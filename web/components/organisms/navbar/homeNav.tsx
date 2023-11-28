import {
    Box,
    Flex,
    Text,
    IconButton,
    Button,
    Stack,
    Collapse,
    Icon,
    Link as Clink,
    Popover,
    PopoverTrigger,
    PopoverContent,
    useColorModeValue,
    useBreakpointValue,
    useDisclosure,
    useColorMode,
    Drawer,
    DrawerOverlay,
    DrawerContent,
    DrawerCloseButton,
    DrawerHeader,
    DrawerBody,
    Tooltip,
    Container,
    VStack,
    Divider,
    HStack,
    Avatar,
} from '@chakra-ui/react';
import {
    HamburgerIcon,
    CloseIcon,
    ChevronDownIcon,
    ChevronRightIcon,
    MoonIcon,
    SunIcon
} from '@chakra-ui/icons';

import Link from 'next/link'
import BasicButton from '@/components/atoms/button/basic';
import { FaDoorOpen, FaSignOutAlt, FaUser } from 'react-icons/fa';
import { useEffect, useState } from 'react';
import { RootState } from '@/redux/store';
import { useSelector } from 'react-redux';

interface MainNavbarProps {
    logo?: React.ReactNode;
    showLink?: boolean
    transparent?: boolean
}

export default function HomeNavbar({
    logo = <Text>Logo</Text>,
    showLink = true,
    transparent = false
}: MainNavbarProps) {
    const { isOpen, onToggle } = useDisclosure();
    const { colorMode, toggleColorMode } = useColorMode();
    const auth = useSelector((state: RootState) => state.auth?.auth)

    const [showNavbarLink, setShowNavbarLink] = useState(false)
    const [offset, setOffset] = useState(0);
    const setScroll = () => {
        setOffset(window.scrollY);
    };

    useEffect(() => {
        window.addEventListener("scroll", setScroll);
        return () => {
            window.removeEventListener("scroll", setScroll);
        };
    }, []);

    useEffect(() => {
        if (offset > 50) {
            setShowNavbarLink(true)
        } else {
            setShowNavbarLink(false)
        }
    }, [offset])

    return (
        <Box>
            <Box
                bg={useColorModeValue(transparent ? 'transparent' : 'white', transparent ? 'transparent' : 'gray.800')}
                color={useColorModeValue(transparent ? 'white' : 'gray.600', 'white')}
                top={0}
                minH={'60px'}
                py={{ base: 2 }}
                borderBottom={showLink ? 1 : 0}
                borderStyle={transparent ? 'none' : 'solid'}
                borderColor={useColorModeValue('gray.200', 'gray.500')}
                position={'fixed'}
                w={'100%'}
                zIndex={200}
            >
                <Container maxW={{ xl: 'container.xl', lg: 'container.lg' }}>
                    <Flex alignItems={'center'} flexDirection={{ base: 'row-reverse', md: 'row' }}>
                        <Flex
                            flex={{ base: 1, md: 'auto' }}
                            ml={{ base: -2 }}
                            display={{ base: showNavbarLink ? 'flex' : 'none', md: 'none' }}
                            justify={showNavbarLink ? 'end' : 'center'}
                        >
                            <IconButton
                                onClick={onToggle}
                                icon={
                                    isOpen ? <CloseIcon w={3} h={3} /> : <HamburgerIcon w={5} h={5} />
                                }
                                variant={'ghost'}
                                aria-label={'Toggle Navigation'}
                                display={showLink ? 'inherit' : 'none'}
                            />
                        </Flex>
                        <Flex transition={'all .3s ease'} flex={{ base: 1 }} justify={{ base: showNavbarLink ? 'start' : 'center', md: 'start' }}>
                            <Link href={'/'} passHref>
                                {logo}
                            </Link>

                            <Flex display={{ base: 'none', md: showLink ? 'flex' : 'none' }} ml={10} alignItems={'center'} >
                                <DesktopNav color={transparent ? 'white' : ''} />
                            </Flex>
                        </Flex>

                        <Flex display={{ md: 'inline-flex', base: 'none' }} columnGap={2} justify={'flex-end'} flex={{ base: 1, md: 0 }}>
                            {auth.user.name !== "" ? (
                                <>
                                    <Popover trigger="hover" placement={'auto'}>
                                        <PopoverTrigger>
                                            <HStack
                                                spacing={'2'}
                                                padding={1}
                                                rounded={'md'}
                                                _hover={{
                                                    cursor: 'pointer',
                                                    backgroundColor: 'gray.100'
                                                }}
                                            >
                                                {/* <Image width={10} height={10} rounded={'full'} objectFit={'cover'} src="https://fastly.picsum.photos/id/64/200/200.jpg?hmac=lJVbDn4h2axxkM72s1w8X1nQxUS3y7li49cyg0tQBZU" alt="profile" /> */}
                                                <Avatar name={auth.user.name} />
                                                <Flex flexDirection={'column'} alignItems={'start'} gap={0} display={{ base: 'none', md: 'inline-flex' }}>
                                                    <Text fontWeight={'medium'} marginBottom={0}>{auth.user.name}</Text>
                                                    {/* <Text fontSize={'xs'} color={'gray.500'}>@ada123</Text> */}
                                                </Flex>
                                            </HStack>
                                        </PopoverTrigger>
                                        <PopoverContent maxWidth={'40'}>
                                            <Box paddingX={'4'} paddingY={'5'}>
                                                <VStack divider={<Divider />} alignItems={'start'}>
                                                    <Button
                                                        as={Link}
                                                        href={'/profile'}
                                                        variant={'link'}
                                                        leftIcon={<FaUser />}
                                                    >
                                                        Profil
                                                    </Button>
                                                    <Button onClick={() => {
                                                        localStorage.removeItem('access');
                                                        localStorage.removeItem('refresh')
                                                        window.location.href = '/auth/login';
                                                    }}
                                                        variant={'link'}
                                                        leftIcon={<FaSignOutAlt />}
                                                    >
                                                        Keluar
                                                    </Button>
                                                </VStack>
                                            </Box>
                                        </PopoverContent>
                                    </Popover>
                                </>
                            ) : (
                                <>
                                    <Tooltip label="Masuk Akun">
                                        <Button
                                            as={Clink}
                                            fontWeight={400}
                                            colorScheme={transparent ? 'gray' : 'green'}
                                            variant={'outline'}
                                            borderWidth={'2px'}
                                            href={'/auth/login'}
                                            _hover={{
                                                textDecoration: 'none',
                                                backgroundColor: 'green.600',
                                                color: 'white'
                                            }}
                                            display={{ base: showLink ? 'inherit' : 'none', md: 'inherit' }}
                                        >
                                            <Text>Masuk</Text>
                                        </Button>
                                    </Tooltip>
                                    <Tooltip label="Buat Akun">

                                        <Button
                                            as={Clink}
                                            fontWeight={400}
                                            backgroundColor={transparent ? 'yellow.400' : 'green.500'}
                                            variant={'solid'}
                                            color={transparent ? 'yellow.800' : 'white'}
                                            href={'/auth/registration'}
                                            _hover={{
                                                textDecoration: 'none',
                                                backgroundColor: 'green.600',
                                                color: 'white'
                                            }}
                                            display={['none', 'none', 'inline-flex']}
                                        >
                                            Daftar
                                        </Button>
                                    </Tooltip>
                                </>
                            )}
                        </Flex>
                    </Flex>
                </Container>
            </Box>

            <Drawer placement={'right'} onClose={onToggle} isOpen={isOpen}>
                <DrawerOverlay />
                <DrawerContent>
                    <DrawerCloseButton />
                    <DrawerHeader borderBottomWidth='1px'>Menu</DrawerHeader>
                    <DrawerBody>
                        <MobileNav />
                    </DrawerBody>
                </DrawerContent>
            </Drawer>
            {/* <Collapse in={isOpen} animateOpacity>
                <MobileNav />
            </Collapse> */}
        </Box>
    );
}

const DesktopNav = ({
    color
}: {
    color?: string
}) => {
    const linkColor = (typeof color !== 'undefined' && color !== "") ? color : 'gray.800'
    const linkHoverColor = useColorModeValue('green.500', 'white');
    const popoverContentBgColor = useColorModeValue('white', 'gray.800');

    return (
        <Stack direction={'row'} spacing={4}>
            {NAV_ITEMS.map((navItem) => !navItem.mobileOnly && (
                <Box key={navItem.label}>
                    <Popover trigger={'hover'} placement={'bottom-start'}>
                        <PopoverTrigger>
                            <Clink
                                as={Link}
                                p={2}
                                href={navItem.href ?? '#'}
                                fontSize={'md'}
                                fontWeight={500}
                                color={linkColor}
                                _hover={{
                                    textDecoration: 'none',
                                    color: linkHoverColor,
                                }}>
                                {navItem.label}
                            </Clink>
                        </PopoverTrigger>

                        {navItem.children && (
                            <PopoverContent
                                border={0}
                                boxShadow={'xl'}
                                bg={popoverContentBgColor}
                                p={4}
                                rounded={'xl'}
                                minW={'sm'}>
                                <Stack>
                                    {navItem.children.map((child) => (
                                        <DesktopSubNav key={child.label} {...child} />
                                    ))}
                                </Stack>
                            </PopoverContent>
                        )}
                    </Popover>
                </Box>
            ))}
        </Stack>
    );
};

const DesktopSubNav = ({ label, href, subLabel }: NavItem) => {
    return (
        <Clink
            as={Link}
            href={href}
            role={'group'}
            display={'block'}
            p={2}
            rounded={'md'}
            _hover={{ bg: useColorModeValue('pink.50', 'gray.900') }}>
            <Stack direction={'row'} align={'center'}>
                <Box>
                    <Text
                        transition={'all .3s ease'}
                        _groupHover={{ color: 'pink.400' }}
                        fontWeight={500}>
                        {label}
                    </Text>
                    <Text fontSize={'sm'}>{subLabel}</Text>
                </Box>
                <Flex
                    transition={'all .3s ease'}
                    transform={'translateX(-10px)'}
                    opacity={0}
                    _groupHover={{ opacity: '100%', transform: 'translateX(0)' }}
                    justify={'flex-end'}
                    align={'center'}
                    flex={1}>
                    <Icon color={'pink.400'} w={5} h={5} as={ChevronRightIcon} />
                </Flex>
            </Stack>
        </Clink>
    );
};

const MobileNav = () => {
    return (
        <Stack
            bg={useColorModeValue('white', 'gray.800')}
            p={2}
            display={{ md: 'none' }}>
            {NAV_ITEMS.map((navItem) => (
                <MobileNavItem key={navItem.label} {...navItem} />
            ))}
        </Stack>
    );
};

const MobileNavItem = ({ label, children, href, mobileOnly }: NavItem) => {
    const { isOpen, onToggle } = useDisclosure();

    return (
        <Stack spacing={4} onClick={children && onToggle}>
            <Flex
                py={2}
                px={2}
                as={Clink}
                href={href ?? '#'}
                justify={'space-between'}
                align={'center'}
                rounded={'lg'}
                backgroundColor={mobileOnly ? 'green.400' : 'white'}
                _hover={{
                    textDecoration: 'none',
                    backgroundColor: mobileOnly ? 'green.600' : 'gray.200'
                }}>
                <Text
                    fontWeight={600}
                    color={mobileOnly ? 'white' : 'gray.600'}>
                    {label}
                </Text>
                {children && (
                    <Icon
                        as={ChevronDownIcon}
                        transition={'all .25s ease-in-out'}
                        transform={isOpen ? 'rotate(180deg)' : ''}
                        w={6}
                        h={6}
                    />
                )}
            </Flex>

            <Collapse in={isOpen} animateOpacity style={{ marginTop: '0!important' }}>
                <Stack
                    mt={2}
                    pl={4}
                    borderLeft={1}
                    borderStyle={'solid'}
                    borderColor={useColorModeValue('gray.200', 'gray.700')}
                    align={'start'}>
                    {children &&
                        children.map((child) => (
                            <Clink key={child.label} py={2} href={child.href}>
                                {child.label}
                            </Clink>
                        ))}
                </Stack>
            </Collapse>
        </Stack>
    );
};

interface NavItem {
    label: string;
    subLabel?: string;
    children?: Array<NavItem>;
    href?: string;
    mobileOnly?: boolean;
}

const NAV_ITEMS: Array<NavItem> = [
    {
        label: 'Produk',
        href: '/product'
    },
    {
        label: 'Masuk',
        href: '/auth/login',
        mobileOnly: true
    },
    {
        label: 'Daftar',
        href: '/auth/registration',
        mobileOnly: true
    }
];
