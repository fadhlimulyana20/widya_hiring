import Header from "@/components/atoms/head";
import MainTemplate from "@/components/template/main";
import MainOldTemplate from "@/components/template/mainold";
import { backendApiURL } from "@/constant/urls/backend_api";
import { ResetPassword } from "@/models/user";
import { backendAPI } from "@/utils/axios";
import { HandleErrorAxios } from "@/utils/axios/error";
import { ViewIcon, ViewOffIcon } from "@chakra-ui/icons";
import { Box, Button, Container, Flex, FormControl, FormErrorMessage, FormLabel, Grid, GridItem, Heading, Input, InputGroup, InputRightElement, Stack, Text, useToast } from "@chakra-ui/react";
import { Form, Formik, FormikHelpers } from "formik";
import { useRouter } from "next/router";
import { useState } from "react";
import { FaPaperPlane, FaSave } from "react-icons/fa";

interface ResetPasswodValue {
    password: string;
    confirm_password: string;
}

function EmailForm() {
    const [showPassword, setShowPassword] = useState(false)
    const [showConfirmPassword, setShowConfirmPassword] = useState(false)
    const router = useRouter()
    const toast = useToast()

    return (
        <Formik
            initialValues={{ password: '', confirm_password: '' }}
            validate={(values: ResetPasswodValue) => {
                let errors: any = {}

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
            onSubmit={async (values, { setSubmitting }: FormikHelpers<ResetPasswodValue>) => {
                const { id } = router.query
                const token = String(id)
                setSubmitting(true)
                try {
                    const res = await ResetPassword({ token: token, password: values.password })
                    if ((typeof res?.code !== 'undefined') && ([200, 201].includes(res?.code))) {
                        toast({
                            title: 'Password berhasil diubah',
                            status: 'success',
                            duration: 9000,
                            isClosable: true,
                        })
                        window.location.href = '/auth/login'
                    }
                } catch (e: any) {
                    HandleErrorAxios({ e: e, title: 'Password gagal diubah', toast: toast })
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
                <form onSubmit={handleSubmit}>

                    <Stack spacing={5}>
                        <Heading fontSize={'xl'}>Buat Password</Heading>
                        <FormControl>
                            <FormLabel>Password Baru</FormLabel>
                            <InputGroup>
                                <Input value={values.password} name={'password'} onChange={handleChange} type={showPassword ? 'text' : 'password'} placeholder="Masukan Password Baru" />
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
                            <FormErrorMessage>{errors.password && touched.password && errors.password}</FormErrorMessage>
                        </FormControl>
                        <FormControl>
                            <FormLabel>Konfirmasi Password Baru</FormLabel>
                            <InputGroup>
                                <Input value={values.confirm_password} name={'confirm_password'} onChange={handleChange} type={showConfirmPassword ? 'text' : 'password'} placeholder="Konfirmasi Password Baru" />
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
                            <FormErrorMessage>{errors.confirm_password && touched.confirm_password && errors.confirm_password}</FormErrorMessage>
                        </FormControl>
                        <Button type={'submit'} leftIcon={<FaSave />} colorScheme="green" isDisabled={isSubmitting}>{ isSubmitting ? 'Menyimpan...' : 'Simpan' }</Button>
                    </Stack>
                </form>
            )}
        </Formik>

    )
}

export default function ForgotPasswordIndex() {
    return (
        <>
            <Header
                title="Reset Password"
                description="Ubah passwordmu untuk mengakses kembali akunmu"
            />
            <MainOldTemplate>
                <Box backgroundColor={'green.100'}>
                    <Container as={Flex} maxW={'container.xl'} minH={'100vh'} paddingTop={'10'} alignItems={'center'}>
                        <Grid w={'100%'} templateColumns={['repeat(1, 1fr)', 'repeat(1, 1fr)', 'repeat(2, 1fr)']} gap={5} alignItems={'center'}>
                            <Box as={GridItem}>
                                <Heading color={'green.800'}>Reset Password</Heading>
                                <Text color={'green.700'}>Habis direset, jangan lupa lagi ya passwordnya...</Text>
                            </Box>
                            <Box as={GridItem} backgroundColor={'white'} padding={'5'} rounded={'lg'}>
                                <EmailForm />
                            </Box>
                        </Grid>
                    </Container>
                </Box>
            </MainOldTemplate>
        </>
    )
}
