import { createContext, useState, useContext, useEffect } from "react";
import { WebsocketContext } from "../_services/websocketProvider";

const NotificationCountContext = createContext()

export const NotificationCountProvider = ({children}) => {
    const [notifCount, setNotifCount] = useState(0)

    const [ready, value, send, connect] = useContext(WebsocketContext);

    useEffect(() => {
        if (!ready) {
            setNotifCount(0)
            return
        }
  
        if (value === null) {
            return
        }
  
        const data = JSON.parse(value?.data)
  
        if (value?.type === "notif") {
            setNotifCount(data.count)
        }
  
        if (value?.type === "newnotif") {
            setNotifCount(notifCount + 1)
        }
    }, [ready, value])

    return <NotificationCountContext.Provider value={{notifCount, setNotifCount}}>
        {children}
    </NotificationCountContext.Provider>
}

export default NotificationCountContext