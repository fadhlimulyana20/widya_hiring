import { Flex, Link as CLink, Box, Text, VStack } from "@chakra-ui/react";
import Link from 'next/link'
import { useEffect, useState } from "react";

interface hoverColor {
    bg: string;
    color: string
}

export function NavItem({
    active = false,
    href = "#",
    title = "title",
    icon = true,
    showText = true,
    iconSize = 'lg',
    colorScheme = 'gray',
    textPosition = 'inline'
}: {
    active?: boolean
    href?: string
    title?: string
    icon?: any
    showText?: boolean
    iconSize?: string
    colorScheme?: string | "gray" | "green"
    textPosition?: "bottom" | "inline"
}) {
    const [hoverColor, setHoverColor] = useState<hoverColor>({
        bg: '',
        color: ''
    })

    useEffect(() => {
        switch (colorScheme) {
            case "gray":
                setHoverColor({
                    bg: 'gray.200',
                    color: 'gray.600'
                })
                break;
            case "green":
                setHoverColor({
                    bg: 'green.200',
                    color: 'green.600'
                })
            default:
                break;
        }
    }, [])


    return (
        <CLink as={Link} href={href} style={{ textDecoration: 'none' }} _focus={{ boxShadow: 'none' }}>
            <Flex
                align={'center'}
                alignItems={'center'}
                padding={'2'}
                fontSize={'lg'}
                fontWeight={'semibold'}
                color={active ? 'green.500' : 'gray.600'}
                rounded={'lg'}
                _hover={active ? {} : {
                    bg:  hoverColor.bg,
                    border: '0',
                    color: hoverColor.color
                }}
                backgroundColor={active ? 'green.100' : 'transparent'}
                border={active ? '2px' : '0'}
                borderColor={active ? 'green.300' : 'transparent'}
            >
                {textPosition === "inline" ? (
                    <>
                        <Box marginRight={showText ? '2' : '0'}>
                            <Text fontSize={iconSize}>{icon}</Text>
                        </Box>
                        <Box display={showText ? 'inline' : 'none'}>
                            {title}
                        </Box>
                    </>
                ) : (
                    <VStack gap={0}>
                        <Text fontSize={iconSize}>{icon}</Text>
                        <Text mt={"0 !important"} pt={0} fontSize={'sm'}> {title}</Text>
                    </VStack>
                )}
            </Flex>
        </CLink>
    )
}
