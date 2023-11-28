import { Box, Flex, Heading, Image, useToast } from "@chakra-ui/react";
import BasicButton from "../atoms/button/basic";
import { LogoText } from "../atoms/logo";
import MainFooter from "../organisms/footer/main";
import MainNavbar from "../organisms/navbar/main";
import HomeNavbar from "../organisms/navbar/homeNav";
import { useCallback, useEffect, useRef, useState } from "react";
import { useDispatch, useSelector } from "react-redux";
import { GetAuthUser } from "@/models/user";
import { setAuth } from "@/redux/slices/authSlices";
import { RootState } from "@/redux/store";
import { refreshToken } from "@/utils/axios";

interface props {
    children: React.ReactNode
}

function Logo({
  textColor="green.500"
}: {
  textColor?: string
}) {
    return (
        <>
            <Flex alignItems={'center'} columnGap={'1'}>
                <Image src="/logo/logo_sm_128.png" w={'auto'} h={'14'} alt={'logo-navbar'} />
                <Heading fontSize={'2xl'} color={textColor}>Logo</Heading>
            </Flex>
        </>
    )
}

export default function MainTemplate({ children }: props) {
    const [showNavbarLink, setShowNavbarLink] = useState(false)
    const [offset, setOffset] = useState(0);
    const toast = useToast()
    const [authLoaded, setAuthLoaded] = useState(false)
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

    const auth = useSelector((state: RootState) => state.auth?.auth)
    const dispatch = useDispatch()

    const getUser = useCallback(async () => {
        try {
            const res = await GetAuthUser()
            const timeout = localStorage.getItem('timeout') || ''
            if ((typeof res?.code == 'number') && ([200, 201].includes(res?.code))) {
                dispatch(setAuth({ ...auth, user: res.data, token: { ...auth.token, timeout: timeout } }))
                // setAuthUser(res.data)
                setAuthLoaded(true)
            }
        } catch (e: any) {
            toast({
                title: 'Sesi habis',
                status: 'error',
                duration: 9000,
                isClosable: true,
            })
            // window.location.href = '/auth/login'
        }
    }, [])


    useEffect(() => {
      if (auth.user.name === '') {
          if (!authLoaded) {
              getUser()
          }
      } else {
          const t = auth.token.timeout
          if (t !== null) {
              const t_time = new Date(t)
              const now = new Date()
              if (t_time < now) {
                  // console.log('oops')
                  const token = localStorage.getItem('token')
                  refreshToken(token || "")
              }
          }
      }

  }, [])


    return (
        <>
            <HomeNavbar transparent={!showNavbarLink} showLink={true} logo={<Logo textColor={!showNavbarLink ? 'white' : 'green.500'} />} />
            <Box>
                {children}
            </Box>
            <MainFooter />
        </>
    )
}
