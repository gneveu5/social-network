import { useEffect, useState } from "react"
import './../css/contactList.css'
import './../css/alertCounter.css'

const ContactList = (props) => {
    const [searchValue, setSearchValue] = useState("")
    const [filteredPrivate, setFilteredPrivate] = useState([])
    const [filteredGroup, setFilteredGroup] = useState([])

    useEffect(() => {
        if (searchValue === "") {
            setFilteredGroup(props.contacts.group)
            setFilteredPrivate(props.contacts.private)
        } else {
            const lowerCaseSearchValue = searchValue.toLowerCase()
            setFilteredGroup(props.contacts.group.filter((contact) => contact.name.toLowerCase().includes(lowerCaseSearchValue)))
            setFilteredPrivate(props.contacts.private.filter((contact) => contact.name.toLowerCase().includes(lowerCaseSearchValue)))
        }
    }, [searchValue, props.contacts])

    return (    
    <div id="contact-window">
        <ContactWindowHeader/>
        <div id="contact-list">
            <input type='text' value={searchValue} onChange={(e) => {setSearchValue(e.target.value)}}/>
            <div className="contact-category">Contacts</div>
            {filteredPrivate.map((contact) => <ContactComponent key={"pvt-" + contact.id} changeConversation={props.changeConversation} contact={contact}/>)}
            <div className="contact-category">Groups</div>
            {filteredGroup.map((contact) => <ContactComponent key={"grp-" + contact.id} changeConversation={props.changeConversation} contact={contact}/>)}
        </div>
    </div>
    )
}

const ContactWindowHeader = () => {
    const [hideContactList, setHideContactList] = useState(true)

    const hide = () => {
        if (hideContactList) {
            document.getElementById('contact-list').style.display = 'none'
        } else {
            document.getElementById('contact-list').style.display = 'flex'
        }
        setHideContactList(!hideContactList)
    }

    return (
        <>
        <div id="contact-window-top">Chat</div>
        <div id="hide-contact-list-button" onClick={hide}>{hideContactList?"v":"^"}</div>
        </>
        )
}

const ContactComponent = (props) => {
    const name = props.contact.name

    const unreadMessages = props.contact.unreadMessages < 100 ? props.contact.unreadMessages : 99

    const changeConversation = () => {
        props.changeConversation(props.contact)
        document.getElementById('chat-window').classList = ""
        document.getElementById('chat-window-informations').innerText = props.contact.name
    }

    if (props.contact.type === "pvt") {
        return (
        <div className={props.contact.connected? "contact connected":"contact disconnected"} onClick={changeConversation}>
            {name}
            {unreadMessages > 0 ? <div className="alert-counter">{unreadMessages}</div>:null}
        </div>
        )
    } else {
        return (
        <div className="contact" onClick={changeConversation}>
            {name}
            {unreadMessages > 0 ? <div className="alert-counter">{unreadMessages}</div>:null}
        </div>
        )
    }
}

export default ContactList