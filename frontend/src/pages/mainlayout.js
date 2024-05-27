import LeftMenu from './leftMenu'
import ChatMain from '../chat/chatMain'
import "../css/mainlayout.css"
import { Outlet } from 'react-router-dom'
import { WebsocketProvider } from '../_services/websocketProvider'
import { NotificationCountProvider } from '../context/notificationCountContext'

const MainLayout = () => {
    return (
        <>
        <WebsocketProvider>
            <NotificationCountProvider>
            <LeftMenu/>
            <div className='maintlayout'>
            <Outlet/>
            </div>
            <ChatMain/>
            </NotificationCountProvider>
        </WebsocketProvider>
        </>
    )
}

export default MainLayout