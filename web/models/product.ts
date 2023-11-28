import { ApiResponse } from "@/constant/api";
import { backendApiURL } from "@/constant/urls/backend_api";
import { backendAPI } from "@/utils/axios";
import { buildQuery } from "@/utils/query";
import { AxiosResponse } from "axios";

export interface ProductResponse {
    id: number;
    name: string;
    description: string;
    created_at: string;
    updated_at: string;
}

export interface ProductCreate {
    name: string;
    description: string;
}

export interface ProductUpdate extends ProductCreate {
    id: number;
}

export interface ProductFilterParam {
    q?: string;
    page?: number;
    limit?: number;
}


export async function GetProductList({
    q="",
    page=1,
    limit=10
} : ProductFilterParam) {
    let param = ''
    param = buildQuery({ q, limit, page })

    try {
        const res = await backendAPI.get<ApiResponse<Array<ProductResponse>>>(backendApiURL.basic.product.base+ '?' + param)
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}

export async function CreateProduct({
    name,
    description
}: ProductCreate) {
    try {
        const res = await backendAPI.post<AxiosResponse<ProductResponse>>(backendApiURL.basic.product.base, {
            name,
            description
        })
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}

export async function UpdateProduct({
    id,
    name,
    description
}: ProductUpdate) {
    try {
        const res = await backendAPI.put<AxiosResponse<ProductResponse>>(backendApiURL.basic.product.base + `/${id}`, {
            name,
            description
        })
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}

export async function DeleteProduct(id: number) {
    try {
        const res = await backendAPI.delete<AxiosResponse<ProductResponse>>(backendApiURL.basic.product.base + `/${id}`)
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}


export async function GetProduct(id: number) {
    try {
        const res = await backendAPI.get<AxiosResponse<ProductResponse>>(backendApiURL.basic.product.base + `/${id}`)
        if ([200, 201].includes(res.status)) {
            return res.data
        }
    } catch (e: any) {
        throw (e)
    }
}

