package carrier

import (
	"errors"
	"fmt"
	"github.com/googollee/go-socket.io"
)

func checkSocketAuthorization(ns socketio.Socket) string {
	if socketSession := this.redis.HGet(REDIS_USER_CARRIER_KEY, ns.Id()).Val(); socketSession == "" {
		ns.Emit("disconnect")
		return fmt.Sprintf("(Carrier) Disconnecting session %s due to auth timeout", ns.Id())
	}
	return ""
}

func setSocketAuthorization(ns socketio.Socket, user *User) {
	defer user.SetOnline()

	uid := string(user.ID)

	pipeline := this.redis.Pipeline()
	pipeline.HSet(REDIS_USER_CARRIER_KEY, ns.Id(), uid)
	pipeline.HSet(REDIS_USER_SOCKET_SESSION_KEY, uid, this.id)

	pipeline.Exec()

	if _, found := UsersMap[user.ID]; !found {
		UsersMap[user.ID] = make(map[socketio.Socket]bool)
	}

	SocketsMap[ns] = int(user.ID)
	UsersMap[user.ID][ns] = true
}

func removeSocketAuthorization(ns socketio.Socket) string {
	user, _ := FindUserBySocket(ns)
	defer delete(SocketsMap, ns)

	if user != nil {
		defer delete(UsersMap[user.ID], ns)
		defer user.SetOffline()

		pipeline := this.redis.Pipeline()

		pipeline.HDel(REDIS_USER_SOCKET_SESSION_KEY, string(user.ID))
		pipeline.HDel(REDIS_USER_CARRIER_KEY, ns.Id())

		_, err := pipeline.Exec()
		if err != nil {
			return fmt.Sprintf("(Carrier) Unauthorizing user %d unsuccessful on session %s: %s", user.ID, ns.Id(), err)
		}
	}
	return ""
}

func FindUserBySocket(ns socketio.Socket) (*User, error) {
	user := new(User)
	if _, found := SocketsMap[ns]; !found {
		return user, errors.New(fmt.Sprintf("(UserBySocket) User not found for socket session (%s)", ns.Id()))
	}

	query := this.db.Find(user, SocketsMap[ns])
	if query.Error != nil {
		return user, errors.New(fmt.Sprintf("(UserBySocket) User not found in database (%s)", query.Error))
	}
	return user, nil
}

func FindSocketByUserId(user_id int) (map[socketio.Socket]bool, error) {
	if sockets, found := UsersMap[user_id]; !found {
		return sockets, errors.New(fmt.Sprintf("(SocketByUser) Not found for uid (%d)", user_id))
	} else {
		return sockets, nil
	}
}
