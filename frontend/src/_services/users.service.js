import Axios from "./api";


let UserInformation = (credentials) => {
    return Axios.get('/userinformation?query='+ credentials)
}

export const usersService = {
    UserInformation

}
