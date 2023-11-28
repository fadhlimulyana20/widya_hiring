import Header from "@/components/atoms/head";
import MainTemplate from "@/components/template/main";
import MainOldTemplate from "@/components/template/mainold";
import { backendApiURL } from "@/constant/urls/backend_api";
import { backendAPI } from "@/utils/axios";
import { Box, Button, Container, Flex, FormControl, Grid, GridItem, Heading, Input, Stack, Text, useToast } from "@chakra-ui/react";
import { Form, Formik, FormikHelpers } from "formik";
import { FaPaperPlane } from "react-icons/fa";

interface ResetEmailValue {
    email: string
}

function EmailForm() {
    const toast = useToast()

    return (
        <Formik
            initialValues={{ email: '' }}
            validate={(values: ResetEmailValue) => {
                let errors: any = {}

                if (!values.email) {
                    errors.email = "harus diisi"
                }

                return errors
            }}
            onSubmit={async (values, { setSubmitting }: FormikHelpers<ResetEmailValue>) => {
                setSubmitting(true)
                try {
                    const res = await backendAPI.post(backendApiURL.public.auth.resetPassword.request, { ...values })
                    if ([200, 201].includes(res.status)) {
                        toast({
                            title: 'Sukses',
                            description: 'Permintaan reset password berhasil dikirim ke email mu.',
                            status: 'success',
                            duration: 9000,
                            isClosable: true,
                        })
                    }
                } catch (e: any) {
                    toast({
                        title: 'Permintaan Reset Password Gagal',
                        status: 'error',
                        duration: 9000,
                        isClosable: true,
                    })
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
                        <Heading fontSize={'xl'}>Kirim Link Reset Password</Heading>
                        <FormControl>
                            <Input value={values.email} name="email" onChange={handleChange} placeholder="Masukan Email" />
                        </FormControl>
                        <Button type="submit" leftIcon={<FaPaperPlane />} colorScheme="green" isDisabled={isSubmitting}>
                            {isSubmitting ? 'Mengirim Link...' : 'Kirim'}
                        </Button>
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
                title="Lupa Password"
                description="Kirim permintaan reset password untuk mengubah password mu"
            />
            <MainOldTemplate>
                <Box backgroundColor={'green.100'}>
                    <Container as={Flex} maxW={'container.xl'} minH={'100vh'} paddingTop={'10'} alignItems={'center'}>
                        <Grid w={'100%'} templateColumns={['repeat(1, 1fr)', 'repeat(1, 1fr)', 'repeat(2, 1fr)']} gap={5} alignItems={'center'}>
                            <Box as={GridItem} w={'100%'}>
                                <Heading color={'green.800'}>Lupa Password</Heading>
                                <Text color={'green.700'}>Gapapa lupa, wajar kok, manusia emang bisa lupa. Masukan emailmu untuk mengirimkan link reset password.</Text>
                            </Box>
                            <Box w={'100%'} as={GridItem} backgroundColor={'white'} padding={'5'} rounded={'lg'}>
                                <EmailForm />
                            </Box>
                        </Grid>
                    </Container>
                </Box>
            </MainOldTemplate>
        </>
    )
}
