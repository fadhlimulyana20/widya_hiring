import Header from "@/components/atoms/head";
import { RequestConfirmationEmail } from "@/models/user";
import { HandleErrorAxios } from "@/utils/axios/error";
import { Box, Button, Flex, Heading, Input, Stack, Text, useToast } from "@chakra-ui/react";
import Link from "next/link";
import { useCallback, useState } from "react";

export default function RequestEmailConfirmation() {
    const [email, setEmail] = useState('')
    const [processing, setProcessing] = useState(false)
    const toast = useToast()

    const sendConfirmationLink = useCallback(async (email: string) => {
        setProcessing(true)
        try {
            const res = await RequestConfirmationEmail(email)
            if (res) {
                toast({
                    title: 'Link berhasil dikirim',
                    description: 'Cek email mu untuk mengaktivasi akun',
                    status: 'success',
                    duration: 2000,
                    isClosable: true,
                })
            }
        } catch(e: any) {
            HandleErrorAxios({e: e, title: 'Link Gagal dikirim', toast: toast})
        } finally {
            setProcessing(false)
        }
    }, [])

    return (
        <>
            <Header
                title="Permintaan konfirmasi Email"
                description="Permintaan konfirmasi Email"
            />
            <Box minH={'100vh'} backgroundColor={'green.100'}>
                <Flex minH={'100vh'} alignItems={'center'} justify={'center'}>
                    <Box as={Stack} textAlign={'center'} spacing={'5'} backgroundColor={'white'} rounded={'lg'} padding={4}>
                        <Heading color={'green.800'}>Permintaan Konfirmasi Email</Heading>
                        <Input value={email} onChange={(e: any) => setEmail(e.target.value)} placeholder="Masukan Email" />
                        <Button onClick={() => sendConfirmationLink(email)} isDisabled={processing} size={'sm'} colorScheme="green">{processing ? 'Mengirim...' : 'Kirim Sekarang'}</Button>
                    </Box>
                </Flex>
            </Box>
        </>
    )
}
