import Header from "@/components/atoms/head";
import MyModal from "@/components/molecules/modal";
import PaginationPure from "@/components/molecules/pagination/pure";
import TableKey from "@/components/molecules/table";
import MainOldTemplate from "@/components/template/mainold";
import { Meta } from "@/constant/api";
import { DeleteProduct, GetProductList, ProductFilterParam, ProductResponse } from "@/models/product";
import { HandleErrorAxios } from "@/utils/axios/error";
import { Box, Button, Container, Heading, Stack, Text, Toast, useToast } from "@chakra-ui/react";
import Link from "next/link";
import { useRouter } from "next/router";
import { Dispatch, SetStateAction, useCallback, useEffect, useState } from "react";

function ProductDeleteConfirm({
    id,
    modalOpen,
    setModalOpen,
    listData,
    setListData
}: {
    id: number;
    modalOpen: boolean;
    setModalOpen: Dispatch<SetStateAction<boolean>>;
    listData: Array<ProductResponse>;
    setListData: Dispatch<SetStateAction<Array<ProductResponse>>>;
}) {
    const [isSubmitting, setIsSubmitting] = useState(false)
    const toast = useToast()

    return (
        <>
            <MyModal
                id={'modal-delete-product'}
                isOpen={modalOpen}
                title="Hapus Product"
                onClose={() => setModalOpen(false)}
                withSaveButton
                saveText="Ya"
                onSave={async () => {
                    setIsSubmitting(true)
                    try {
                        const res = await DeleteProduct(id)
                        if (typeof res !== 'undefined') {
                            toast({
                                title: 'Berhasil',
                                description: 'Produk berhasil dihapus',
                                status: 'success',
                                duration: 2000,
                                isClosable: true,
                            })

                            let d = listData.filter((obj, idx) => obj.id !== id)
                            setListData(d)
                        }
                    } catch(e: any) {
                        HandleErrorAxios({e, title: 'Gagal menghapus produk', toast})
                    } finally {
                        setIsSubmitting(false)
                        setModalOpen(false)
                    }
                }}
                isSaving={isSubmitting}
            >
                <Text>Aapakah anda yakin menghapus produk ini?</Text>
            </MyModal>
        </>
    )
}

export default function ProductManagement() {
    const [openModalDelete, setOpenModalDelete] = useState(false)
    const [productIDSelected, setProductIDSelected] = useState(0)
    const [productList, setProductList] = useState<Array<ProductResponse>>([])
    const [productListMeta, setProductListMeta] = useState<Meta>({
        limit: 10,
        page: 1,
        total_count: 0,
        total_page: 0
    })
    const [isLoading, setIsLoading] = useState(false)
    const toast = useToast()
    const router = useRouter()

    const fetchData = useCallback(async ({limit, page}: Meta, q: string) => {
        console.log(productListMeta)
        try {
            const res = await GetProductList({
                limit: limit,
                page: page,
                q: q
            })

            if (typeof res !== 'undefined') {
                setProductList(res.data)
                if (typeof res.meta !== 'undefined') {
                    setProductListMeta(res.meta)
                }
            }
        } catch (e: any) {
            HandleErrorAxios({ e, title: 'Terjadi error', toast })
        }
    }, [])

    useEffect(() => {
        if (router.isReady) {
            const {page} = router.query
            if (typeof page !== 'undefined') {
                setProductListMeta({...productListMeta, page: Number(page)})
                fetchData({...productListMeta, page: Number(page)}, '')
            } else {
                fetchData(productListMeta, '')
            }

        }
    }, [router])



    const tableOptions = [
        {
            name: 'Nama',
            accessor: 'name',
        },
        {
            name: 'Diperbarui',
            accessor: 'updated_at',
            type: 'date'
        }
    ]

    const tableActions = [
        {
            name: 'Edit',
            action: (id: any) => {
                router.push(`/product/edit/${id}`)
            }
        },
        {
            name: 'Hapus',
            action: (id: any) => {
                setProductIDSelected(id)
                setOpenModalDelete(true)
            }
        }
    ]

    return (
        <>
            <Header
                title="Manajemen Produk"
                description="Manajemen Produk"
            />

            <ProductDeleteConfirm
                id={productIDSelected}
                listData={productList}
                modalOpen={openModalDelete}
                setListData={setProductList}
                setModalOpen={setOpenModalDelete}
                key={'modal'}
            />

            <MainOldTemplate>
                <Box minH={'100vh'} paddingY={0} paddingBottom={20}>
                    <Stack spacing={'20'}>
                        <Container mt={40} maxW={{ xl: 'container.lg', lg: 'container.md' }}>
                            <Heading mb={5}>Daftar Produk</Heading>
                            <Button as={Link} href={'/product/create'} mb={4} colorScheme="green">Tambah</Button>
                            {isLoading ? (
                                <Text>Memuat...</Text>
                            ) : (
                                <Stack spacing={4}>
                                    <TableKey data={productList} options={tableOptions} actions={tableActions} />
                                    <PaginationPure
                                        onChangePage={(page: number) => {
                                            // dispatch(setCreatorQuestionPackPage(page))
                                            router.push({query: {page: page}}, undefined, {shallow: true})
                                        }}
                                        page={productListMeta.total_page}
                                        activePage={productListMeta.page}
                                        colorScheme="green"
                                        key={'pagination'}
                                    />
                                </Stack>
                            )}
                        </Container>
                    </Stack>
                </Box>
                {/* <div id="g_id_onload" data-client_id="783583331881-i6e016v331s24pv757jrng89cqbt3j59.apps.googleusercontent.com" data-type="standard" data-size="small" data-callback="handleCredentialResponse"></div>
        <div className="g_id_signin" data-client_id="783583331881-i6e016v331s24pv757jrng89cqbt3j59.apps.googleusercontent.com" data-type="standard" ></div> */}
            </MainOldTemplate>
        </>
    )
}
