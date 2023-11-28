import Header from "@/components/atoms/head";
import MainOldTemplate from "@/components/template/mainold";
import { RootState } from "@/redux/store";
import { Box, Container, Divider, Flex, Heading, Stack, Text } from "@chakra-ui/react";
import { useSelector } from "react-redux";

export default function UserProfilePage() {
    const auth = useSelector((state: RootState) => state.auth?.auth)

    return (
        <>
            <Header
                title="Profil"
                description="Profil"
            />
            <MainOldTemplate>
                <Box minH={'100vh'} paddingY={0} paddingBottom={20}>
                    <Stack spacing={'20'}>
                        <Container mt={40} maxW={{ xl: 'container.lg', lg: 'container.md' }}>
                            <Heading mb={5}>Profil</Heading>
                            <Stack divider={<Divider />}>
                                <Box>
                                    <Text>Nama</Text>
                                    <Heading fontSize={'lg'}>{auth.user.name}</Heading>
                                </Box>
                                <Box>
                                    <Text>Email</Text>
                                    <Heading fontSize={'lg'}>{auth.user.email}</Heading>
                                </Box>
                            </Stack>
                        </Container>
                    </Stack>
                </Box>
            </MainOldTemplate>
        </>
    )
}
