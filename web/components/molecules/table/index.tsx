import { Button, HStack, IconButton, Popover, PopoverArrow, PopoverBody, PopoverCloseButton, PopoverContent, PopoverHeader, PopoverTrigger, Table, TableContainer, Tbody, Td, Th, Thead, Tr, useColorModeValue, VStack } from "@chakra-ui/react";
import Router from "next/router";
import { useState } from "react";
import { FaCheck, FaTimes } from "react-icons/fa";
import { FiEye, FiMoreVertical, FiTrash } from "react-icons/fi";
import MyModal from "../modal";
import Moment from "react-moment";
import { Latex } from "@/components/atoms/latex";

interface option {
    name: string
    accessor: string | Function
    customAccessor?: Function
    withLink?: boolean
    link?: string | Function
    type?: 'date' | 'string' | 'boolean' | 'latex' | any
    truncate?: boolean
}

export interface ActionTableType {
    name: string;
    action: (id: any) => any
}

export default function TableKey({
    data,
    options,
    onDetail,
    showAction = true,
    rowWithActionClick = false,
    rowActionClick = () => {},
    actions = [],
    onDelete = (id) => null
}: {
    data: Array<any>,
    options: Array<option>,
    showDeleteButton?: boolean,
    showAction?: boolean,
    rowWithActionClick?: boolean,
    rowActionClick?: Function
    actions?: Array<ActionTableType>,
    onDetail?: (id: any) => any,
    onDelete?: (id: any) => any
}) {
    const [modalDeleteOpen, setModalDeleteOpen] = useState(false)
    const [idToDelete, setIdToDelete] = useState(0)
    const [isDeleting, setIsDeleting] = useState(false)
    const bgHover = useColorModeValue('green.50', 'green.100')
    const colorHover = useColorModeValue('black', 'green.700')

    return (
        <>
            <MyModal
                title="Hapus Kunci"
                isOpen={modalDeleteOpen}
                onClose={() => setModalDeleteOpen(false)}
                onSave={async () => {
                    setIsDeleting(true)
                    await onDelete(idToDelete)
                    setIsDeleting(false)
                    setModalDeleteOpen(false)
                }}
                saveText='Ya'
                isSaving={isDeleting}
                cancelText="Batal"
            >
                Hapus kunci ini?
            </MyModal>
            <TableContainer>
                <Table variant='simple' rounded={'md'}>
                    <Thead bg={useColorModeValue('gray.50', 'gray.800')}>
                        <Tr>
                            {/* <Th>No.</Th> */}
                            {options.map((o, idx) => (<Th key={idx}>{o.name}</Th>))}
                            <Th display={showAction ? '' : 'none'}>Pilihan</Th>
                        </Tr>
                    </Thead>
                    <Tbody>
                        {data.map((d, idx) => (
                            <Tr
                                key={idx}
                                _hover={{ background: bgHover, color: colorHover, cursor: rowWithActionClick ? 'pointer' : 'normal' }}
                                onClick={() => {
                                    rowWithActionClick && rowActionClick(d.id)
                                }}
                            >
                                {/* <Td>{idx + 1}</Td> */}
                                {options.map((o, oidx) => {
                                    if (typeof o.accessor === 'string') {
                                        if (o.type && o.type === 'date') {
                                            return (<Td key={oidx}>{new Date(eval(`d.${o.accessor}`)).toDateString()}</Td>)
                                        } else if (o.type && o.type === 'boolean') {
                                            return (<Td key={oidx}>{eval(`d.${o.accessor}`) === true ? <FaCheck color="green" /> : <FaTimes color="red" />}</Td>)
                                        } else if (o.type && o.type === 'latex') {
                                            if (o.truncate) {
                                                return (
                                                    <Td maxWidth={'72'} style={{ wordWrap: 'break-word', whiteSpace: 'pre-wrap' }} key={oidx}>
                                                        <Latex text={eval(`d.${o.accessor}`)} />
                                                    </Td>
                                                )
                                            } else {
                                                return (
                                                    <Td key={oidx}>
                                                        <Latex text={eval(`d.${o.accessor}`)} />
                                                    </Td>
                                                )
                                            }
                                        } else if (o.type && o.type === 'custom') {
                                            return (<Td key={oidx}>{ typeof o.customAccessor !== 'undefined' && o.customAccessor(eval(`d.${o.accessor}`)) }</Td>)
                                        }

                                        if (o.truncate) {
                                            return (<Td maxWidth={'72'} style={{ wordWrap: 'break-word', whiteSpace: 'pre-wrap' }} key={oidx}>{eval(`d.${o.accessor}`)}</Td>)
                                        }

                                        return (<Td key={oidx}>{eval(`d.${o.accessor}`)}</Td>)
                                    } else {
                                        return (<Td key={oidx}>{ }</Td>)
                                    }
                                })}
                                <Td display={showAction ? '' : 'none'}>
                                    <Popover placement='left-start' preventOverflow size={'sm'}>
                                        <PopoverTrigger>
                                            <IconButton
                                                size="sm"
                                                variant="ghost"
                                                colorScheme={'gray'}
                                                aria-label="Detail"
                                                icon={<FiMoreVertical />}
                                            />
                                        </PopoverTrigger>
                                        <PopoverContent width={'max-content'}>
                                            {/* <PopoverHeader fontWeight='semibold'>Pilihan</PopoverHeader> */}
                                            <PopoverArrow />
                                            <PopoverBody>
                                                <VStack alignItems={'start'}>
                                                    {actions.map((obj, idx) => (
                                                        <Button
                                                            key={idx}
                                                            variant={'ghost'}
                                                            size='sm'
                                                            onClick={() => obj.action(d.id)}
                                                            w={'100%'}
                                                        >
                                                            {obj.name}
                                                        </Button>
                                                    ))}
                                                </VStack>
                                            </PopoverBody>
                                        </PopoverContent>
                                    </Popover>
                                    <HStack>

                                        {/* {showDeleteButton && (
                                            <IconButton
                                                size="sm"
                                                variant="ghost"
                                                colorScheme={'red'}
                                                aria-label="Delete"
                                                icon={<FiTrash />}
                                                onClick={() => {
                                                    setIdToDelete(d.id)
                                                    setModalDeleteOpen(true)
                                                }}
                                            />
                                        )} */}
                                    </HStack>
                                </Td>
                            </Tr>
                        ))}
                    </Tbody>
                </Table>
            </TableContainer>
        </>
    )
}
