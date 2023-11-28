import axios from "axios";

export const blogAPI = axios.create({
    baseURL: process.env.NEXT_PUBLIC_BLOG_URL
})

blogAPI.interceptors.request.use(function (config) {
    config.params = {...config.params, key: process.env.NEXT_PUBLIC_BLOG_API_KEY}

    return config
}, function (error) {
    return Promise.reject(error)
})
