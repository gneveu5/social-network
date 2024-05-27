import Axios from "./api"

const getNotifications = () => {
    return Axios.get('/notifications')
}

const notificationReponse = (credentials) => {
    return Axios.post('/notificationresponse', credentials)
}

export const NotificationsService = {
    getNotifications, notificationReponse
}