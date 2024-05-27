package socialnetwork

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	auth "socialnetwork/pkg/api/auth"
	fol "socialnetwork/pkg/api/follow"
	group "socialnetwork/pkg/api/group"
	notif "socialnetwork/pkg/api/notifications"
	post "socialnetwork/pkg/api/post"
	search "socialnetwork/pkg/api/searchbar"
	user "socialnetwork/pkg/api/users"
	ws "socialnetwork/pkg/api/websocket"

	// db "socialnetwork/pkg/db"
	database "socialnetwork/pkg/db/sqlite"
)

func Runserver() {
	middleware := NestedMiddleware(AllowOrigin, Checksession) // liste des fonctions middlewares à passer

	var err error
	database.Db, err = sql.Open("sqlite3", "pkg/db/social.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Db.Close()
	database.Db.Exec("PRAGMA foreign_keys = ON")
	if len(os.Args) > 1 {
		if os.Args[1] == "migrateDown" {
			database.MigrateDown()
		}
	} else {
		database.MigrateUp()
	}

	// post
	http.Handle("/post", middleware(post.PostHandler))
	http.Handle("/comment", middleware(post.CommentHandler))
	http.Handle("/commenting", middleware(post.CommentingHandler))

	// group
	http.Handle("/group", middleware(group.GroupHandler))
	http.Handle("/grouplist", middleware(group.GroupListHandler))
	http.Handle("/groupcreation", middleware(group.GroupCreationHandler))
	http.Handle("/groupfetchexist", middleware(group.GroupFetchExistHandler))
	http.Handle("/groupaskregister", middleware(group.GroupAskMembershipHandler))
	http.Handle("/groupaskmembership", middleware(group.GroupAskMembershipHandler))
	http.Handle("/groupinvitemembership", middleware(group.GroupInviteMembershipHandler))
	http.Handle("/groupmember", middleware(group.GroupFetchMemberHandler))
	http.Handle("/groupmembers", middleware(group.GroupFetchMembersHandler))
	http.Handle("/groupfetchaskmembership", middleware(group.GroupFetchAskMembershipHandler))
	http.Handle("/event", middleware(group.EventAttendeeFetchHandler))

	// auth
	http.Handle("/login", AllowOrigin(auth.LoginHandler))
	http.Handle("/register", AllowOrigin(auth.RegisterHandler))
	http.Handle("/checktoken", middleware(auth.VeriftokenHandler))

	// searchbar
	http.Handle("/searchbar", middleware(search.SearchHandler))

	// chat / notif???
	http.Handle("/socket", AllowOrigin(ws.SocketHandler))

	// follow
	http.Handle("/follower", middleware(fol.FollowerHandler))
	http.Handle("/followernotingroup", middleware(fol.FollowerNotInGroupHandler))
	http.Handle("/following", middleware(fol.FollowingHandler))
	http.Handle("/reqfollow", middleware(fol.RequestFollowHandler))
	http.Handle("/requnfollow", middleware(fol.RequnfollowHandler))
	http.Handle("/reqcancelfollowprivate", middleware(fol.ReqCancelFollowPrivateHandler))
	http.Handle("/notificationresponse", middleware(notif.NotificationResponseHandler))

	// user
	http.Handle("/userinformation", middleware(user.UserinfoHandler))
	http.Handle("/privatetopublic", middleware(user.PrivatetopublicHandler))
	http.Handle("/publictoprivate", middleware(user.PublictoprivateHandler))
	http.Handle("/nickname", middleware(user.NicknameHandler))

	// image folder
	http.Handle("/avatar/", http.StripPrefix("/avatar/", http.FileServer(http.Dir("pkg/db/avatars/"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("pkg/db/img/"))))

	// notifications
	http.Handle("/notifications", middleware(notif.NotificationsHandler))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("server listening on port 8080")
}
