import { LogoText } from "@/components/atoms/logo"
import { Box, Button, Container, Flex, Heading, IconButton, Link as CLink, SimpleGrid, Stack, Text, useColorModeValue, Image } from "@chakra-ui/react"
import Link from "next/link"
import { FaEnvelope, FaInstagram, FaLinkedin, FaTiktok, FaYoutube } from "react-icons/fa"

export default function MainFooter() {
    return (
        <Box
            bg={'green.400'}
            zIndex={200}
        >
            <Box backgroundColor={'green.600'} py={'2'}>
                <Text textAlign={'center'} fontSize={'sm'} color={'white'}>
                    © 2023 Fadhli Mulyana. All rights reserved
                </Text>
            </Box>
        </Box>
    )
}
