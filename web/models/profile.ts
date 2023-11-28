import { ApiResponse } from "@/constant/api";
import { backendAPI } from "@/utils/axios";
import { User } from "./user";
import { backendApiURL } from "@/constant/urls/backend_api";

export interface updateProfileParam {
    name: string;
    email: string;
}

export interface UpdatePasswordParam {
    old_password: string;
    new_password: string;
}

export async function UpdateProfile({ name, email }: updateProfileParam) {
    try {
        const res = await backendAPI.post<ApiResponse<User>>(backendApiURL.basic.auth.updateAccount, {
            name,
            email
        })
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}

export async function UpdatePassword({ old_password, new_password }: UpdatePasswordParam) {
    try {
        const res = await backendAPI.post<ApiResponse<User>>(backendApiURL.basic.auth.updatePassword, {
            old_password,
            new_password
        })
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}

