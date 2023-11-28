import Header from "@/components/atoms/head"
import MainOldTemplate from "@/components/template/mainold"
import { CreateProduct, ProductCreate } from "@/models/product"
import { HandleErrorAxios } from "@/utils/axios/error"
import { Box, Button, Container, FormControl, FormLabel, Heading, Input, Stack, Textarea, useToast } from "@chakra-ui/react"
import { useRouter } from "next/router"
import { ChangeEvent, useState } from "react"
import { fileURLToPath } from "url"

export default function CreateProductPage() {
    const [createProductData, setCreateProductData] = useState<ProductCreate>({
        description: '',
        name: ''
    })
    const [isSaving, setIsSaving] = useState(false)
    const toast = useToast()
    const router = useRouter()

    const handleCreate = async () => {
        setIsSaving(true)
        try {
            const res = await CreateProduct(createProductData)
            if (typeof res?.data !== 'undefined') {
                toast({
                    title: 'Produk Berhasil Ditambahkan',
                    status: 'success',
                    duration: 2000,
                    isClosable: true,
                })
                router.replace(`/product/edit/${res.data.id}`)
            }
        } catch(e: any) {
            HandleErrorAxios({e, title: 'Gagal Menambah Produk', toast})
        } finally {
            setIsSaving(false)
        }
    }

    return (
        <>
            <Header
                title="Menambah Produk"
                description="Menambah Produk"
            />
            <MainOldTemplate>
                <Box minH={'100vh'} paddingY={0} paddingBottom={20}>
                    <Stack spacing={'20'}>
                        <Container mt={40} maxW={{ xl: 'container.lg', lg: 'container.md' }}>
                            <Heading mb={5}>Tambah Produk</Heading>
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
