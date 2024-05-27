import { useEffect, useState, useContext } from "react"
import { NotificationsService } from "../_services/notifications.service"
import { useNavigate } from "react-router-dom"
import NotificationCountContext from "../context/notificationCountContext"

import "../css/notifications.css"




const Notifications = () => {
    const [response, setResponse] = useState({})

    const {notifCount, setNotifCount} = useContext(NotificationCountContext)

    const getData = async () => {
        try {
            const res = await NotificationsService.getNotifications()
            setResponse(res.data)
            setNotifCount(0)
            console.log(res.data)
        } catch (error) {
            setResponse({error:error})
        }

    }

    useEffect(() => {
        getData()
    }, [])

    return (
        <div className="notifications">
        <h1>Notifications</h1>
        <button onClick={getData}>Update</button>
        <h2>Follow requests</h2>
        <div>{response.followRequest?.map((notif) => <FollowRequestNotif key={notif.notification.id} notif={notif}/>)}</div>
        <h2>Group membership requests</h2>
        <div>{response.groupMembershipRequest?.map((notif) => <GroupMembershipRequestNotif key={notif.notification.id} notif={notif}/>)}</div>
        <h2>Group invites</h2>
        <div>{response.groupInvite?.map((notif) => <GroupInviteNotif key={notif.notification.id} notif={notif} />)}</div>
        <h2>Group events</h2>
        <div>{response.event?.map((notif) => <EventNotif key={notif.notification.id} notif={notif} />)}</div>
        </div>
    )
}

const FollowRequestNotif = (props) => {
    const navigate = useNavigate()

    return (
        <div>
            <div onClick={() => {navigate("/user/" + props.notif.userName)}}>User {props.notif.userName} wants to follow you.</div>
            <NotificationResponse notificationId={props.notif.notification.id}/>
        </div>
    )
}

const GroupMembershipRequestNotif = (props) => {
    return (
        <div>
            <div>User {props.notif.userName} wants to join group {props.notif.groupName}.</div>
            <NotificationResponse notificationId={props.notif.notification.id}/>
        </div>
    )
}

const EventNotif = (props) => {
    return (
        <div>
            <div>An event has been created on group {props.notif.groupName} : {props.notif.eventName}. Will you participate ?</div>
            <NotificationResponse notificationId={props.notif.notification.id}/>
        </div>
    )
}

const GroupInviteNotif = (props) => {
    return (
        <div>
            <div>User {props.notif.userName} invited you to join group {props.notif.groupName}.</div>
            <NotificationResponse notificationId={props.notif.notification.id}/>
        </div>
    )
}

const NotificationResponse = (props) => {
    const [responseStatus, setResponseStatus] = useState(0)

    const accept = () => {
        const response = {
            notificationId:props.notificationId,
            confirm:true
        }
        NotificationsService.notificationReponse(response).then(
            (r) => {setResponseStatus(1)}
        ).catch(
            (e) => {setResponseStatus(3)}
        )
    }

    const refuse = () => {
        const response = {
            notificationId:props.notificationId,
            confirm:false
        }
        NotificationsService.notificationReponse(response).then(
            (r) => {setResponseStatus(2)}
        ).catch(
            (e) => {setResponseStatus(3)}
        )
    }

    switch (responseStatus) {
    case 1:
        return <div>Accepted</div>
    case 2:
        return<div>Refused</div>
    default: 
        return (
            <>
            <button onClick={accept}>Accept</button>
            <button onClick={refuse}>Refuse</button>
            {responseStatus.current === 3 ? 
                <div>Something went wrong... Try again.</div>
                : null
            }
            </>
        )
    }
}

export default Notifications