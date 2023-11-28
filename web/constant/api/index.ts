export interface ApiResponse<T> {
    code: number
    message: string
    errors: Array<any>
    data: T
    meta?: Meta
}

export interface Meta {
    page: number
    limit: number
    total_page: number
    total_count: number
}
