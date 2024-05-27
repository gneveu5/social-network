import { createContext, useState, useEffect, useRef } from 'react';
import { accountService } from "./account.service";

export const WebsocketContext = createContext(false, null, () => {}, () => {})
//                                            ready, value, send, connect

// websocket access context. autoamtically tries to reconnect when disconnected
export const WebsocketProvider = ({ children }) => {
    const [isReady, setIsReady] = useState(0) // websocket is ready to be used
    const [val, setVal] = useState(null)          // last value returned by websocket

    const websocket = useRef(null)                // current websocket connection

    const connect = () => {
        if (!accountService.isLogged()) {
            return
        }

        if (websocket.current != null) {
            return
        }

        // const socket = new WebSocket("ws://localhost:8080/socket", [accountService.getToken().replaceAll('"', '')])
        const socket = new WebSocket("ws://localhost:8080/socket")
        websocket.current = socket

        socket.onopen = () => {
            setIsReady(true)
            socket.send(accountService.getToken().replaceAll('"', ''))
        }
        
        socket.onclose = () => {
            websocket.current = null
            setIsReady(false)
        }

        socket.onmessage = (e) => {
            setVal(JSON.parse(e.data))
        }

        socket.onerror = (e) => {
            console.log("Websocket error")
            console.log(e)
        }
    }

    // on component creation : try to connect websocket
    useEffect(() => {
        connect()
        return () => {
            if (websocket.current) {
                websocket.current.close()
            }
        }
    }, [])

    // value returned by context
    const ret = [isReady, val, websocket.current?.send.bind(websocket.current), connect]

    return (
        <WebsocketContext.Provider value={ret}>
            {children}
        </WebsocketContext.Provider>
    )
}
