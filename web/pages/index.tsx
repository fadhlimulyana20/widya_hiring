import MainTemplate from "@/components/template/main";
import { Box, Button, Container, Flex, Grid, GridItem, Heading, Image as CImage, Stack, Text, Link as Clink, Wrap, WrapItem, SimpleGrid, VStack } from "@chakra-ui/react";
import Head from "next/head";
import * as jose from 'jose';
import { useEffect, useState } from "react";
import Header from "@/components/atoms/head";
import MainOldTemplate from "@/components/template/mainold";
import Link from "next/link";


export default function Index() {
  const handleCredentialResponse = (response: any) => {
    // decodeJwtResponse() is a custom function defined by you
    // to decode the credential response.
    const claims = jose.decodeJwt(response.credential)

    console.log(response)
    console.log(claims)
  }

  useEffect(() => {
    if (typeof window !== "undefined") {
      window.handleCredentialResponse = handleCredentialResponse
    }
  }, [])


  return (
    <>
      <Header
        title="Manajemen Produk"
        description="Manajemen Produk"
      />
      <MainOldTemplate>
        <Box minH={'100vh'} paddingY={0} paddingBottom={20}>
          <Stack spacing={'20'}>
            <Container mt={40} maxW={{ xl: 'container.lg', lg: 'container.md' }}>
              <Heading mb={5}>Halo Selamat Datang</Heading>
              <Flex columnGap={2}>
                <Button colorScheme="green" as={Link} href={'/auth/login'}>Masuk Akun</Button>
                <Button as={Link} href={'/auth/registration'}>Buat Akun</Button>
              </Flex>
            </Container>
          </Stack>
        </Box>
      </MainOldTemplate>
    </>
  )
}
