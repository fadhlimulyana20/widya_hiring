import { Button, ButtonGroup, Card, CardBody, CardFooter, Divider, Heading, Image, Stack } from "@chakra-ui/react";
import React from "react";

interface props {
    children: React.ReactNode
    childrenFooter?: React.ReactNode
    withImage?: boolean
    withFooter?: boolean
    imageSrc?: string
    title?: string
}

export default function BasicCard({
    children,
    childrenFooter = <></>,
    withImage = false,
    withFooter = false,
    imageSrc = 'https://images.unsplash.com/photo-1555041469-a586c61ea9bc?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1770&q=80',
    title = 'Card Title'
}: props) {
    return (
        <Card maxW={'sm'}>
            <CardBody>
                {withImage && (
                    <Image
                        src={imageSrc}
                        alt='Green double couch with wooden legs'
                        borderRadius='lg'
                    />
                )}
                <Stack mt='6' spacing='3'>
                    <Heading size={'md'}>{title}</Heading>
                    {children}
                </Stack>
            </CardBody>
            {withFooter && (
                <>
                    <Divider />
                    <CardFooter>
                        {childrenFooter}
                    </CardFooter>
                </>
            )}
        </Card>
    )
}