import { Box, Flex, Heading, Spinner } from "@chakra-ui/react";

export default function Loading() {
    return (
        <Flex alignItems={'center'} justifyContent={'center'} minH={'100vh'} backgroundColor={'green.100'}>
            <Box textAlign={'center'}>
                <Heading marginBottom={'4'} color={'green.500'}>Kuadran</Heading>
                <Spinner
                    thickness='4px'
                    speed='0.65s'
                    emptyColor='gray.100'
                    color='green.500'
                    size='xl'
                />
            </Box>
        </Flex>
    )
}