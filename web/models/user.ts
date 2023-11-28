import { ApiResponse } from "@/constant/api";
import { backendApiURL } from "@/constant/urls/backend_api";
import { backendAPI } from "@/utils/axios";

export interface Roles {
    name: string;
}

export interface User {
    name: string;
    email: string;
    created_at: string;
    roles: Array<Roles>;
}

export interface Token {
    refresh: string;
    access: string;
    timeout: string;
}

export interface Auth {
    token: Token;
    user: User;
}

export async function GetAuthUser() {
    try {
        const res = await backendAPI.get<ApiResponse<User>>(backendApiURL.basic.auth.me)
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}

export async function ResetPassword({ token, password }: { token: string, password: string }) {
    try {
        const res = await backendAPI.post<ApiResponse<any>>(backendApiURL.public.auth.resetPassword.update, {
            token,
            password
        })
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}

export async function ConfirmEmail(token: string) {
    try {
        const res = await backendAPI.post<ApiResponse<any>>(backendApiURL.public.auth.emailValidation.validate, {
            token
        })
        if ([200, 201].includes(res.status)) {
            return true
        } else {
            return false
        }
    } catch (e: any) {
        throw (e)
    }
}

export async function RequestConfirmationEmail(email: string) {
    try {
        const res = await backendAPI.post<ApiResponse<any>>(backendApiURL.public.auth.emailValidation.request, {
            email
        })
        if ([200, 201].includes(res.status)) {
            return true
        } else {
            return false
        }
    } catch (e: any) {
        throw (e)
    }
}

export async function AuthWithGoogle(token: string) {
    try {
        const res = await backendAPI.post<ApiResponse<Auth>>(backendApiURL.public.auth.oauth.google, {
            token
        })
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}
