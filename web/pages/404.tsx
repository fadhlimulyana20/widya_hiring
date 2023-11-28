import Header from "@/components/atoms/head";
import { Box, Container, Flex, Heading, Image, Link } from "@chakra-ui/react";

export default function Custom404() {
    return (
        <>
            <Header
                title="404: Halaman tidak ditemukan"
                description="Halaman tidak ditemukan"
            />
            <Flex alignItems={'center'} minH={'100vh'} backgroundColor={'green.100'}>
                <Container>
                    <Link href="https://www.freepik.com/free-vector/hand-drawn-no-data-illustration_49342678.htm">
                        <Image src="/illustration/404.png" alt="not-found" width={'96'} height={'96'} mx={'auto'} />
                    </Link>
                    <Heading textAlign={'center'} color={'green.700'}>Halaman Tidak Ditemukan</Heading>
                </Container>
            </Flex>
        </>
    )
}
