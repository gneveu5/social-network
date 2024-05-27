import Axios from "./api";


let login = (credentials) => {
    return Axios.post('/login', credentials)
}

let register = (credentials) => {
    return Axios.post('/register', credentials, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
    })
}

let getAuthenticatedUser = (userToken) => {
    return Axios.get('/getauthenticateduser', userToken) 

}

let saveToken = (token) => {
    // localStorage.setItem('jwt', token)
    sessionStorage.setItem('jwt', token)
}

let logout = () => {
    // localStorage.removeItem('jwt')
    sessionStorage.removeItem('jwt')
}

let isLogged = () => {
    // let jwt = localStorage.getItem('jwt')
    let jwt = sessionStorage.getItem('jwt')
    if (jwt == "null") {
        return false
    }
  
return !!jwt

} 

let getToken =() => {
    // return localStorage.getItem('jwt')
    return sessionStorage.getItem('jwt')
}


let privatetopublic = () => {
    return Axios.post('/privatetopublic')
}
let publictoprivate = () => {
    return Axios.post('/publictoprivate')
}

let nickname = () => {
    return Axios.get('/nickname')
}

export const accountService = {
    login, saveToken, logout, isLogged, getToken, register, privatetopublic, publictoprivate, nickname, getAuthenticatedUser

}