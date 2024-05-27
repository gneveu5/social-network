import {useEffect, useState, useContext} from "react"

import "../css/alertCounter.css"
import NotificationCountContext from "../context/notificationCountContext"

const NotificationsCounter = (props) => {
    const {notifCount, setNotifCount} = useContext(NotificationCountContext)

    return (
        <>{notifCount === 0 ?
        null
        :
        <span className="alert-counter">{notifCount}</span>
    }
    </>
    )
}

export default NotificationsCounter