import Header from "@/components/atoms/head"
import Loading from "@/components/organisms/loading"
import { ConfirmEmail } from "@/models/user"
import { HandleErrorAxios } from "@/utils/axios/error"
import { Box, Button, Flex, Heading, Link, Stack, Text, useToast } from "@chakra-ui/react"
import { useRouter } from "next/router"
import { useCallback, useEffect, useState } from "react"

export default function EmailConfirmationIndex() {
    const [success, setSuccess] = useState(true)
    const [confirmed, setConfirmed] = useState(false)
    const [processing, setProcessing] = useState(true)
    const router = useRouter()
    const toast = useToast()

    const confirmEmail = useCallback(async (token: string) => {
        setProcessing(true)
        try {
            const res = await ConfirmEmail(token)
            if (res) {
                setSuccess(true)
                setConfirmed(true)
            } else {
                setSuccess(false)
            }
        } catch (e: any) {
            HandleErrorAxios({ e: e, title: 'Email gagal diverifiaksi', toast: toast })
            if (e.isAxiosError) {
                const resp = e.response?.data
                if (Array.isArray(resp?.errors)) {
                    resp?.errors.forEach((obj: any) => {
                        if (obj === 'token is completed') {
                            setConfirmed(true)
                        }
                    })
                }
            }
            setSuccess(false)
        } finally {
            setProcessing(false)
        }
    }, [])

    useEffect(() => {
        const { id } = router.query
        if (typeof id === 'string') {
            confirmEmail(id)
        }
    }, [router])


    return (
        <>
            <Header
                title="Konfirmasi Email"
                description="Konfirmasi Email mu melalui link yang dikirmkan melalui email"
            />
            {processing ? (
                <Loading />
            ) : success ? (
                <Box minH={'100vh'} backgroundColor={'green.100'}>
                    <Flex minH={'100vh'} alignItems={'center'} justify={'center'}>
                        <Box as={Stack} textAlign={'center'} spacing={'5'}>
                            <Box>
                                <Heading color={'green.800'}>Emailmu Sudah Terkonfirmasi</Heading>
                                <Text fontSize={'lg'} color={'green.600'}>Kamu bisa login ke akunmu sekarang.</Text>
                            </Box>
                            <Button as={Link} href="/auth/login" size={'sm'} colorScheme="green">Login Sekarang</Button>
                        </Box>
                    </Flex>
                </Box>
            ) : (
                <Box minH={'100vh'} backgroundColor={'green.100'}>
                    <Flex minH={'100vh'} alignItems={'center'} justify={'center'}>
                        <Box as={Stack} textAlign={'center'} spacing={'5'}>
                            <Box>
                                <Heading color={'green.800'}>{ confirmed ? 'Emailmu Sudah Terkonfirmasi' : 'Emailmu Gagal Terkonfirmasi' }</Heading>
                                <Text fontSize={'lg'} color={'green.600'}>{ confirmed ? 'Kamu bisa login ke akunmu sekarang.' : 'Lakukan permintaan konfirmasi email kembali' }</Text>
                            </Box>
                            { confirmed ? (
                                <Button as={Link} href="/auth/login" size={'sm'} colorScheme="green">Login Sekarang</Button>
                            ) : (
                                <Button as={Link} href="/auth/email-confirmation/request" size={'sm'} colorScheme="green">Kirim Link Konfirmasi Email</Button>
                            )}
                        </Box>
                    </Flex>
                </Box>
            )}
        </>
    )
}
