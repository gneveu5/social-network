import { useContext } from "react";
import { WebsocketContext } from "../_services/websocketProvider";

const ReconnectWebsocket = (props) => {
    const [ready, value, send, connect] = useContext(WebsocketContext);

    return (
        <>
        {!ready===true?
            <div id={props.containingDivId}>
                <div>Disconnected...</div>
                <button onClick={connect}>Reconnect</button>
            </div>
            :
            <>
                {props.children}
            </>
        }
        </>
    )
}

export default ReconnectWebsocket