import { ApiResponse } from "@/constant/api";
import { backendApiURL } from "@/constant/urls/backend_api";
import axios from "axios";
import { toast } from "react-toastify";

export const backendAPI = axios.create({
    baseURL: process.env.NEXT_PUBLIC_BACKEND_URL
})

export async function refreshToken(token: string) {
    try {
        const res = await backendAPI.post<ApiResponse<any>>(backendApiURL.public.auth.refresh, {
            refresh: token
        })

        if (res.data.data.token.access) {
            localStorage.setItem('access', res.data.data.token.access)
        }
        return
    } catch(err) {
        // toast('Sesi telah habis', {type: 'error'})
        localStorage.removeItem('refresh')
        localStorage.removeItem('access')
        localStorage.removeItem("current_user")
        window.location.href = '/auth/login'
        return
    }
}

export async function getUser() {
    try {
        const res = await backendAPI.get<ApiResponse<any>>(backendApiURL.basic.auth.me)
        return res.data.data
    } catch(err: any) {
        throw new Error(err)
    }
}

backendAPI.interceptors.request.use(function (config) {
    // Intercept Auth Bearer Token
    const token = localStorage.getItem('access')
    if (token !== null) {
        config.headers['Authorization'] = `Bearer ${token}`
    }

    // // Get Authenticated User Data
    // const user = localStorage.getItem('user')
    // if (user === null) {
    //     getUser()
    // }

    return config
}, function (error) {
    return Promise.reject(error)
})

backendAPI.interceptors.response.use(function (response) {
    return response;
}, function (error) {
    if (401 === error.response.status) {
        const refresh = localStorage.getItem('refresh')
        if (refresh !== null) {
            refreshToken(refresh)
            return Promise.reject(error)
        }

        localStorage.removeItem("refresh")
        localStorage.removeItem("access")
        // localStorage.removeItem('creator')

        if (window.location.pathname !== "/auth/login") {
            window.location.href = '/auth/login'
        }

        return Promise.reject(error)
    } else if (403 == error.response.status) {
        window.location.href = '/error/403'
        return Promise.reject(error)
    // } else if (500 == error.response.status) {
    //     window.location.href = '/error/500'
    //     return Promise.reject(error)
    } else {
        return Promise.reject(error);
    }
});
