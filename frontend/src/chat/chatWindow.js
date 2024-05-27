import './../css/chatWindow.css'
import { useEffect, useState } from 'react'
import { replaceEmojiShortcuts } from '../_services/emojis'

const ChatWindow = (props) => {
    const [typingMessage, setTypingMessage] = useState("")

    const filterTypingMessage = (input) => {
        let filteredValue = input
        filteredValue = filteredValue.replaceAll('\n', '').replaceAll('\t', '')
        setTypingMessage(replaceEmojiShortcuts(filteredValue)) 
    }

    useEffect(() => {
        if (!props.conversation || props.conversation.messageList.length === 0) {
            return
        }
        const messageList = props.conversation.messageList
        if (props.scrollToBottom.current) {
            const div = document.getElementById("chat-window-message-"+messageList[messageList.length - 1].id)
            div.scrollIntoView(false)
        } else {
            const div = document.getElementById("chat-window-from-"+messageList[0].id)
            div.scrollIntoView(true)
        }
    }, [props.conversation?.messageList])

    const hideChat = () => {
        document.getElementById('chat-window').classList = 'hidden'
        props.changeConversation(null)
    }

    const requestHistory = () => {
        if (!props.conversation || props.conversation.messageList.length < 10) {
            return
        }
        const data = {type:"hist", last:props.conversation.messageList[0].id}
        props.send(data)
    }

    const sendMessage = () => {
        if (typingMessage === "") {
            return
        }
        const data = {type:"msg", txt:typingMessage}
        props.send(data)
        setTypingMessage("")
    }

    const sendMessageWithEnter = (e) => {
        if (e.key === 'Enter') {
            sendMessage()
        }
    }

    return (
    <div id="chat-window" className="hidden">
            <div id="chat-window-hide-button" onClick={hideChat}/>
            <div id="chat-window-load-more-button" onClick={requestHistory}/>
            <div id="chat-window-informations"/>
            <div id="messages-container">
                {props.conversation?.messageList.map((message) => <MessageComponent key={message.id} message={message}/>)}
            </div>
            <div id="send-message-container">
                <form>
                    <input type="button" id="send-message-button" onClick={sendMessage}/>
                    <textarea id="send-message-input"
                        value={typingMessage}
                        onKeyUp={sendMessageWithEnter} 
                        onChange={(e) => filterTypingMessage(e.target.value)}
                    />
                </form>
            </div>
        </div>
    )
}

const MessageComponent = (props) => {
    const formatDate = (timestamp) => {
        return new Date(timestamp*1000).toLocaleString()
    }

    return (
        <>
        <div id={"chat-window-from-" + props.message.id} classlist="chat-window-from">
            {"(" + formatDate(props.message.creationTimeStamp) + ") " + props.message.senderName + ":"}
        </div>
        <div id={"chat-window-message-" + props.message.id} classlist="chat-window-message">
            {replaceEmojiShortcuts(props.message.content)}
        </div>
        </>
    )
}

export default ChatWindow