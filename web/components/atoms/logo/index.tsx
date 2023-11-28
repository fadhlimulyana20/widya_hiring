import { Text, useColorModeValue } from "@chakra-ui/react"

interface logoTextProps {
    text: string
    isBgDark?: boolean
    fontSize?: string | number
}

export function LogoText({
    text,
    isBgDark=false,
    fontSize='2xl'
} : logoTextProps) {
    const lightColor = isBgDark ? 'white' : 'gray.900'
    const darkColor = isBgDark ? 'gray.50' : 'gray.50'

    return (
        <Text
            fontSize={fontSize}
            fontWeight={'bold'}
            color={useColorModeValue(lightColor, darkColor)}
        >
            {text}
        </Text>
    )
}
