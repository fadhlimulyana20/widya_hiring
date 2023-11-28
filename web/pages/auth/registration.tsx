import MainTemplate from "@/components/template/main";
import { Box, Button, Container, Divider, Flex, FormControl, FormLabel, Input, InputGroup, InputRightElement, Link as CLink, Stack, Text, useToast } from "@chakra-ui/react";
import Script from "next/script";
import * as jose from 'jose';
import { useEffect, useState } from "react";
import Link from "next/link";
import { Form, Formik, FormikHelpers } from "formik";
import { backendAPI } from "@/utils/axios";
import { backendApiURL } from "@/constant/urls/backend_api";
import { ViewIcon, ViewOffIcon } from "@chakra-ui/icons";
import Header from "@/components/atoms/head";
import { HandleErrorAxios } from "@/utils/axios/error";
import { AuthWithGoogle } from "@/models/user";
import { useRouter } from "next/router";
import { useDispatch } from "react-redux";
import { setAuth } from "@/redux/slices/authSlices";
import MainOldTemplate from "@/components/template/mainold";

declare global {
    interface Window {
        handleCredentialResponse?: any;
    }
}

interface BasicRegistrationValue {
    email: string;
    name: string;
    password: string;
    confirm_password: string;
}

function BasicRegistrationForm() {
    const [showPassword, setShowPassword] = useState(false)
    const [showConfirmPassword, setShowConfirmPassword] = useState(false)
    const toast = useToast()

    return (
        <Formik
            initialValues={{ email: '', name: '', password: '', confirm_password: '' }}
            validate={(values: BasicRegistrationValue) => {
                let errors: any = {}

                if (!values.email) {
                    errors.email = "harus diisi"
                }

                if (!values.name) {
                    errors.name = "harus diisi"
                }

                if (!values.password) {
                    errors.password = "harus diisi"
                }

                if (!values.confirm_password) {
                    errors.confirm_password = "harus diisi"
                }

                if (values.confirm_password !== values.password) {
                    errors.confirm_password = "password harus sama"
                }


                return errors
            }}
            onSubmit={async (values, { setSubmitting }: FormikHelpers<BasicRegistrationValue>) => {
                setSubmitting(true)
                try {
                    const res = await backendAPI.post(backendApiURL.public.auth.registration, { ...values })
                    if ([200, 201].includes(res.status)) {
                        toast({
                            title: 'Registrasi Berhasil.',
                            status: 'success',
                            duration: 9000,
                            isClosable: true,
                        })
                        window.location.href = '/auth/login'
                    }
                } catch (e: any) {
                    console.log(e)
                } finally {
                    setSubmitting(false)
                }
            }}
        >
            {({
                values,
                errors,
                touched,
                handleChange,
                handleBlur,
                handleSubmit,
                isSubmitting,
            }) => (
                <Form onSubmit={handleSubmit}>
                    <Stack spacing={'5'}>

                        <Stack spacing={'2'}>
                            <FormControl>
                                <FormLabel>Email</FormLabel>
                                <Input name="email" onChange={handleChange} onBlur={handleBlur} value={values.email} type={'email'} placeholder="Email" focusBorderColor={'green.500'} />
                            </FormControl>
                            <FormControl>
                                <FormLabel>Nama</FormLabel>
                                <Input name="name" onChange={handleChange} onBlur={handleBlur} value={values.name} type={'text'} placeholder="Nama" focusBorderColor={'green.500'} />
                            </FormControl>
                            <FormControl>
                                <FormLabel>Password</FormLabel>
                                <InputGroup>
                                    <Input name="password" onChange={handleChange} onBlur={handleBlur} value={values.password} type={showPassword ? 'text' : 'password'} placeholder="Password" focusBorderColor={'green.500'} />
                                    <InputRightElement h={'full'}>
                                        <Button
                                            variant={'ghost'}
                                            onClick={() =>
                                                setShowPassword((showPassword) => !showPassword)
                                            }>
                                            {showPassword ? <ViewIcon /> : <ViewOffIcon />}
                                        </Button>
                                    </InputRightElement>
                                </InputGroup>
                            </FormControl>
                            <FormControl>
                                <FormLabel>Konfirmasi Password</FormLabel>
                                <InputGroup>
                                    <Input name="confirm_password" onChange={handleChange} onBlur={handleBlur} value={values.confirm_password} type={showConfirmPassword ? 'text' : 'password'} placeholder="Konfirmasi Password" focusBorderColor={'green.500'} />
                                    <InputRightElement h={'full'}>
                                        <Button
                                            variant={'ghost'}
                                            onClick={() =>
                                                setShowConfirmPassword((showConfirmPassword) => !showConfirmPassword)
                                            }>
                                            {showConfirmPassword ? <ViewIcon /> : <ViewOffIcon />}
                                        </Button>
                                    </InputRightElement>
                                </InputGroup>
                            </FormControl>
                        </Stack>
                        <Button type="submit" colorScheme={'green'} isDisabled={isSubmitting}>{isSubmitting ? 'Memproses...' : 'Registrasi'}</Button>
                    </Stack>
                </Form>
            )}
        </Formik>
    )
}

export default function Registration() {
    const toast = useToast()
    const router = useRouter()
    const dispatch = useDispatch()

    const authWIthGoogle = async (token: string) => {
        try {
            const res = await AuthWithGoogle(token)
            if (typeof res !== 'undefined') {
                toast({
                    title: 'Autentikasi Berhasil.',
                    status: 'success',
                    duration: 2000,
                    isClosable: true,
                })
                localStorage.setItem('access', res.data.token.access)
                localStorage.setItem('refresh', res.data.token.refresh)
                localStorage.setItem('timeout', res.data.token.timeout)
                localStorage.setItem('name', res.data.user.name)
                const { roles } = res.data.user
                if (Array.isArray(roles)) {
                    let admin = roles.findIndex(((obj) => obj?.name === 'admin'))
                    if (admin !== -1) {
                        localStorage.setItem('creator', 'pravda')
                    }
                }
                dispatch(setAuth(res.data))
                router.push('/product')
            }
        } catch (e: any) {
            HandleErrorAxios({ e: e, title: 'Login Gagal', toast: toast })
        }
    }

    const handleCredentialResponse = (response: any) => {
        // decodeJwtResponse() is a custom function defined by you
        // to decode the credential response.
        // const claims = jose.decodeJwt(response.credential)
        authWIthGoogle(response.credential)

        // console.log(response)
        // console.log(claims)
    }
    useEffect(() => {
        if (typeof window !== "undefined") {
            window.handleCredentialResponse = handleCredentialResponse
        }
    }, [])

    return (
        <>
            <Header
                title="Buat Akun"
                description="Buat akun untuk mengakses fitur kuadran dengan melakukan registrasi"
            />

            <MainOldTemplate>
                <Box backgroundColor={'green.100'}>
                    <Container minH={'100vh'} paddingY={'100'}>
                        <Box boxShadow={'md'} rounded={'lg'} backgroundColor={'white'} paddingY={'5'} paddingX={'8'} maxW={'480'} mx={'auto'}>
                            <Stack spacing={'5'}>
                                <Stack spacing={'2'}>
                                    <Text fontSize={'2xl'} fontWeight={'bold'} textAlign={'center'}>Buat Akun</Text>
                                    <Text fontSize={'sm'} textColor={'gray.500'} textAlign={'center'}>Selamat datang di Kuadran, buat akun sekarang untuk mengakses fitur-fitur menarik di Project Quiz.</Text>
                                </Stack>
                                <BasicRegistrationForm />
                                <Box>
                                    <Text fontSize={'sm'} textColor={'gray.500'} textAlign={'center'}>Sudah punya akun? <CLink color={'green'} href={"/auth/login"}>Login sekarang</CLink></Text>
                                </Box>
                            </Stack>
                        </Box>
                    </Container>
                </Box>
            </MainOldTemplate>
        </>
    )
}
