import axios from "axios"
import { accountService } from "./account.service";


const Axios = axios.create({
    baseURL: 'http://localhost:8080'
})


Axios.interceptors.request.use(request => {
    if(accountService.isLogged()) {
        request.headers.Authorization = 'Bearer '+accountService.getToken()
    }
    return request
})

// Intercepteur de réponse API pour vérification de la session
// Axios.interceptors.response.use(response => {
//     return response
// }, error => {
//     if(error.response.status === 401){
//         accountService.logout()
//         window.location = '/login'
//     // } else if (error.response.status === 400) {
//     //     const errorMessage = error.response.data.error;
//     //     console.log(errorMessage)
//     //     window.location = `/error?message=${encodeURIComponent(errorMessage)}`;
//     } else {
//         return Promise.reject(error)
//     }
// })

export default Axios