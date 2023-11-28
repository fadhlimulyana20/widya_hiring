import { range } from "@/utils/range"
import { Button, ChakraTheme, HStack, IconButton, Text } from "@chakra-ui/react"
import { Dispatch, SetStateAction, useCallback, useEffect, useState } from "react"
import { FiChevronsLeft, FiChevronsRight } from "react-icons/fi"

function BtnPage(
    {
        page,
        currentPage,
        setCurrentPage,
        onChangePage,
        btntype = "number",
        btnNext = false,
        maxPage,
        colorScheme
    }:
        {
            page?: number,
            currentPage: number,
            setCurrentPage: Dispatch<SetStateAction<number>>
            onChangePage: Function,
            btntype?: "number" | "page"
            btnNext?: boolean,
            maxPage?: number,
            colorScheme: string
        }
) {
    return (
        <>
            { btntype === "number" && page ? (
                <Button
                    onClick={() => {
                        setCurrentPage(page)
                        onChangePage(page)
                    }}
                    variant={"ghost"}
                    size={'md'}
                    colorScheme={colorScheme}
                    isActive={page === currentPage}
                >
                    {page}
                </Button>
            ) : (
                <IconButton
                    variant={"ghost"}
                    onClick={() => {
                        if (maxPage) {
                            const n = btnNext ? currentPage + 1 : currentPage - 1
                            if (n < 1 || n > maxPage) {
                                return
                            }
                            setCurrentPage(n)
                            onChangePage(n)
                        }
                    }}
                    size={'md'}
                    colorScheme={colorScheme}
                    aria-label={btnNext ? 'next' : 'back'}
                    icon={btnNext ? <FiChevronsRight /> : <FiChevronsLeft />}
                />
            )}
        </>
    )
}


function PaginationPure(
    {
        page,
        activePage= 1,
        onChangePage,
        colorScheme="red"
    }:
        {
            page: number,
            activePage?: number
            onChangePage: (arg: number) => any
            colorScheme?: string
        }
) {
    const [currentPage, setCurrentPage] = useState(activePage)
    const [n, setN] = useState<Array<number>> ([])
    const [nRight, setNRight] = useState<Array<number>>([])
    const [nLeft, setNLeft] = useState<Array<number>>([])

    const handleChangePage = useCallback((currPage: number, page: number) => {
        setCurrentPage(currPage)
        if (page > 5) {
            if (page - currPage <= 3 && page - currPage > 1) {
                setNLeft(range(currPage - 1, currPage + 1, 1))
                setNRight(range(currPage + 2, page + 1, 1))
            } else if (page - currPage <= 1) {
                setNLeft(range(currPage - 3, currPage - 1, 1))
                setNRight(range(currPage - 1, page + 1, 1))
            } else {
                setNLeft(range(currPage, currPage + 3, 1))
                setNRight(range(page - 1, page + 1, 1))
            }
        }
    }, [])

    useEffect(() => {
        setN(Array.from({ length: page }, (_, i) => i + 1))

        return () => {
        }
    }, [page])

    useEffect(() => {
        handleChangePage(currentPage, page)
    }, [currentPage, handleChangePage, page])


    return (
        <HStack spacing={'1'}>
            <BtnPage setCurrentPage={setCurrentPage} colorScheme={colorScheme} onChangePage={onChangePage} currentPage={currentPage} btntype="page" maxPage={page}/>
            {n.length <= 5 ? n.map((p, idx) => (<BtnPage setCurrentPage={setCurrentPage} colorScheme={colorScheme} onChangePage={onChangePage} page={p} currentPage={currentPage} key={idx} maxPage={page} />)) : ''}
            {n.length > 5 ? nLeft.map((p, idx) => (<BtnPage setCurrentPage={setCurrentPage} colorScheme={colorScheme} onChangePage={onChangePage} page={p} currentPage={currentPage} key={idx} maxPage={page} />)) : ''}
            {n.length > 5 ? (<Text color={`${colorScheme}.400`}>...</Text>) : ''}
            {n.length > 5 ? nRight.map((p, idx) => (<BtnPage setCurrentPage={setCurrentPage} colorScheme={colorScheme} onChangePage={onChangePage} page={p} currentPage={currentPage} key={idx} maxPage={page} />)) : ''}
            <BtnPage setCurrentPage={setCurrentPage} colorScheme={colorScheme} onChangePage={onChangePage} currentPage={currentPage} btntype="page" maxPage={page} btnNext/>
        </HStack>
    )
}

export default PaginationPure
