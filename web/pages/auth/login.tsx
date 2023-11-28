import MainTemplate from "@/components/template/main";
import { Box, Button, Container, Divider, Flex, FormControl, FormLabel, Input, InputGroup, InputLeftAddon, InputRightElement, Link as CLink, Stack, Text, useToast } from "@chakra-ui/react";
import Script from "next/script";
import * as jose from 'jose';
import { ButtonHTMLAttributes, useCallback, useEffect, useRef, useState } from "react";
import Link from "next/link";
import { ViewIcon, ViewOffIcon } from "@chakra-ui/icons";
import { FaGoogle } from "react-icons/fa";
import { Form, Formik, FormikHelpers } from "formik";
import { backendApiURL } from "@/constant/urls/backend_api";
import { backendAPI } from "@/utils/axios";
import { useRouter } from "next/router";
import { Auth, AuthWithGoogle, GetAuthUser } from "@/models/user";
import Loading from "@/components/organisms/loading";
import Header from "@/components/atoms/head";
import { HandleErrorAxios } from "@/utils/axios/error";
import { useDispatch, useSelector } from "react-redux";
import { setAuth } from "@/redux/slices/authSlices";
import { ApiResponse } from "@/constant/api";
import { RootState } from "@/redux/store";
import MainOldTemplate from "@/components/template/mainold";

declare global {
    interface Window {
        handleCredentialResponse?: any;
    }
}

interface BasicLoginValue {
    email: string;
    password: string;
}

function BasicLoginForm() {
    const [showPassword, setShowPassword] = useState(false)
    const toast = useToast()
    const router = useRouter()
    const dispatch = useDispatch()

    return (
        <Formik
            initialValues={{ email: '', password: '' }}
            validate={(values: BasicLoginValue) => {
                let errors: any = {}

                if (!values.email) {
                    errors.email = "harus diisi"
                }

                if (!values.password) {
                    errors.email = "harus diisi"
                }

                return errors
            }}
            onSubmit={async (values, { setSubmitting }: FormikHelpers<BasicLoginValue>) => {
                setSubmitting(true)
                try {
                    const res = await backendAPI.post<ApiResponse<Auth>>(backendApiURL.public.auth.login, { ...values })
                    if ([200, 201].includes(res.status)) {
                        toast({
                            title: 'Autentikasi Berhasil.',
                            status: 'success',
                            duration: 2000,
                            isClosable: true,
                        })
                        localStorage.setItem('access', res.data.data.token.access)
                        localStorage.setItem('refresh', res.data.data.token.refresh)
                        localStorage.setItem('timeout', res.data.data.token.timeout)
                        localStorage.setItem('name', res.data.data.user.name)
                        const { roles } = res.data.data.user
                        if (Array.isArray(roles)) {
                            let admin = roles.findIndex(((obj) => obj?.name === 'admin'))
                            let contributor = roles.findIndex(((obj) => obj?.name === 'contributor'))
                            if (admin !== -1) {
                                localStorage.setItem('creator', 'pravda')
                            }
                            if (contributor !== -1) {
                                localStorage.setItem('contributor', 'pravda')
                            }
                        }
                        dispatch(setAuth(res.data.data))
                        router.push('/product')
                    }
                } catch (e: any) {
                    HandleErrorAxios({ e: e, title: 'Login Gagal', toast: toast })
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
                                <Input name="email" onChange={handleChange} onBlur={handleBlur} value={values.email} type={'email'} placeholder="Email" focusBorderColor={'green.500'} />
                            </FormControl>
                            <FormControl>
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
                            <CLink textAlign={'right'} as={Link} color="green" href={'/auth/forgot-password'}>Lupa Password</CLink>
                        </Stack>
                        <Button type="submit" colorScheme={'green'} isDisabled={isSubmitting}>{isSubmitting ? 'Memproses...' : 'Masuk'}</Button>
                    </Stack>
                </Form>
            )}
        </Formik>
    )
}

export default function Login() {
    const [loading, setLoading] = useState(true)
    const router = useRouter()
    const toast = useToast()
    const dispatch = useDispatch()
    const auth = useSelector((state: RootState) => state.auth.auth)

    const getUser = useCallback(async () => {
        try {
            const res = await GetAuthUser()
            if ((typeof res?.code == 'number') && ([200, 201].includes(res?.code))) {
                const timeout = localStorage.getItem('timeout') || ''
                dispatch(setAuth({...auth, user: res.data, token: {...auth.token, timeout: timeout}}))
                router.push('/product')
            }
        } catch (e: any) {
            return
        } finally {
            setLoading(false)
        }
    }, [])

    useEffect(() => {
        if (auth.user.name === '') {
            getUser()
        }
    }, [])


    return (
        <>
            <Header
                title="Masuk Akun"
                description="Masukan email dan password untuk mengakses aplikasi kuadran"
            />
            <Box display={loading ? 'initial' : 'none'}><Loading /></Box>
            <Box display={loading ? 'none' : 'initial'}>

                <MainOldTemplate>
                    <Box backgroundColor={'green.100'}>
                        <Container minH={'100vh'} paddingTop={'100'}>
                            <Box boxShadow={'md'} rounded={'lg'} backgroundColor={'white'} paddingY={'5'} paddingX={'8'} maxW={'480'} mx={'auto'}>
                                <Stack spacing={'5'}>
                                    <Stack spacing={'2'}>
                                        <Text fontSize={'2xl'} fontWeight={'bold'} textAlign={'center'}>Masuk Akun</Text>
                                        <Text fontSize={'sm'} textColor={'gray.500'} textAlign={'center'}>Selamat datang di Kuadran, masukan email dan passwordmu untuk mengakses aplikasi.</Text>
                                    </Stack>
                                    <BasicLoginForm />

                                    <Box>
                                        <Text fontSize={'sm'} textColor={'gray.500'} textAlign={'center'}>Belum punya akun? <CLink color={'green'} href={"/auth/registration"}>Buat sekarang</CLink></Text>
                                    </Box>
                                </Stack>
                            </Box>
                        </Container>
                    </Box>
                </MainOldTemplate>
            </Box>
        </>
    )
}
