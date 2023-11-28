import { Button } from "@chakra-ui/react";
import Link from "next/link";
import { useState } from "react";

interface props {
    onClick?: () => any;
    isDisabled?: boolean;
    size?: 'xs' | 'sm' | 'md' | 'lg';
    colorScheme?: 'main' | 'secondary' | any;
    isLink?: boolean;
    href?: string;
    children: React.ReactNode
}

export default function BasicButton({
    onClick = () => { },
    isDisabled = false,
    size = 'md',
    colorScheme = 'main',
    isLink = false,
    href = "#",
    children
}: props) {
    const colorVariant: any = {
        'main': 'blue',
        'secodary': 'gray'
    }

    if (isLink) {
        return (
            <Button
                as={Link}
                isDisabled={isDisabled}
                size={size}
                colorScheme={colorVariant[colorScheme] !== "" ? colorVariant[colorScheme] : colorScheme}
                rounded={'full'}
                href={href}
            >
                {children}
            </Button>
        )
    }

    return (
        <Button
            onClick={onClick}
            isDisabled={isDisabled}
            size={size}
            colorScheme={colorVariant[colorScheme] !== undefined ? colorVariant[colorScheme] : colorScheme}
            rounded={'full'}
        >
            {children}
        </Button>
    )
}