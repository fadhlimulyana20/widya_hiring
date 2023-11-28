import Header from "@/components/atoms/head"
import MyModal from "@/components/molecules/modal"
import MainOldTemplate from "@/components/template/mainold"
import { CreateProduct, DeleteProduct, GetProduct, ProductCreate, ProductResponse, ProductUpdate, UpdateProduct } from "@/models/product"
import { HandleErrorAxios } from "@/utils/axios/error"
import { Box, Button, Container, FormControl, FormLabel, Heading, Input, Stack, Text, Textarea, useToast } from "@chakra-ui/react"
import { useRouter } from "next/router"
import { ChangeEvent, Dispatch, SetStateAction, useCallback, useEffect, useState } from "react"
import { fileURLToPath } from "url"

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
                <Text>Aapakah anda yakin menghapus produk soal ini?</Text>
            </MyModal>
        </>
    )
}

export default function UpdateProductPage() {
    const [createProductData, setCreateProductData] = useState<ProductUpdate>({
        id: 0,
        description: '',
        name: ''
    })
    const [isSaving, setIsSaving] = useState(false)
    const toast = useToast()
    const router = useRouter()

    const handleCreate = async () => {
        setIsSaving(true)
        try {
            const res = await UpdateProduct(createProductData)
            if (typeof res?.data !== 'undefined') {
                toast({
                    title: 'Produk Berhasil Disimpan',
                    status: 'success',
                    duration: 2000,
                    isClosable: true,
                })
            }
        } catch(e: any) {
            HandleErrorAxios({e, title: 'Gagal Menambah Produk', toast})
        } finally {
            setIsSaving(false)
        }
    }

    const handleLoadData = useCallback(async (id: number) => {
        try {
            const res = await GetProduct(id)
            if (typeof res?.data !== 'undefined') {
                setCreateProductData({...res.data})
            }
        }catch(e: any) {
            HandleErrorAxios({e, title: 'Gagal Mentimpan Produk', toast})
        }
    }, [])

    useEffect(() => {
        if (router.isReady) {
            const {idx} = router.query
            handleLoadData(Number(idx))
        }
    }, [router])


    return (
        <>
            <Header
                title="Mengubah Produk"
                description="Mengubah Produk"
            />
            <MainOldTemplate>
                <Box minH={'100vh'} paddingY={0} paddingBottom={20}>
                    <Stack spacing={'20'}>
                        <Container mt={40} maxW={{ xl: 'container.lg', lg: 'container.md' }}>
                            <Heading mb={5}>Ubah Produk</Heading>
                            <Box p={4} border={'1px'} borderColor={'gray.300'} rounded={'lg'} mb={5}>
                                <Stack spacing={5}>
                                    <FormControl>
                                        <FormLabel>Nama</FormLabel>
                                        <Input
                                            type="text"
                                            placeholder="Nama Produk"
                                            value={createProductData.name}
                                            onChange={(e: ChangeEvent<HTMLInputElement>) => setCreateProductData({...createProductData, name: e.target.value})}
                                        />
                                    </FormControl>
                                    <FormControl>
                                        <FormLabel>Deskripsi</FormLabel>
                                        <Textarea
                                            placeholder="Deskripsi Produk"
                                            rows={5}
                                            value={createProductData.description}
                                            onChange={(e: ChangeEvent<HTMLTextAreaElement>) => setCreateProductData({...createProductData, description: e.target.value})}
                                        >

                                        </Textarea>
                                    </FormControl>
                                </Stack>
                            </Box>
                            <Button
                                colorScheme="green"
                                onClick={() => handleCreate()}
                                isDisabled={isSaving}
                            >
                                { isSaving ? 'Meyimpan...' : 'Simpan' }
                            </Button>
                        </Container>
                    </Stack>
                </Box>
            </MainOldTemplate>
        </>
    )
}
