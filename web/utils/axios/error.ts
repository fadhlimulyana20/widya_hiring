import { ApiResponse } from "@/constant/api"
import { useToast } from "@chakra-ui/react"
import { AxiosError } from "axios"

export function HandleErrorAxios({ e, title, toast }: { e: Error | any, title?: string, toast: any }) {
    const err = e as AxiosError<ApiResponse<any>>
    if (err.isAxiosError) {
        const resp = err.response?.data
        if (Array.isArray(resp?.errors)) {
            resp?.errors.forEach((obj) => {
                toast({
                    title: title || 'Galat',
                    description: obj,
                    status: 'error',
                    duration: 2000,
                    isClosable: true,
                })
            })
        }
    }
}
