import { Button, Modal, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay } from "@chakra-ui/react"
import { ReactNode } from "react"
import { GrClose } from "react-icons/gr";

interface MayModalType {
    id?: string;
    isOpen: boolean;
    children: ReactNode;
    title: string;
    onClose: () => any,
    onSave?: () => any;
    isSaving?: boolean;
    saveText?: string;
    cancelText?: string;
    withSaveButton?: boolean;
}

function MyModal({
    id='modal',
    isOpen,
    onClose,
    children,
    title,
    onSave = () => {},
    isSaving = false,
    saveText = 'Simpan',
    cancelText = 'Tutup',
    withSaveButton = false
}: MayModalType) {
	return (
		<Modal id={id} closeOnOverlayClick={false} isOpen={isOpen} onClose={onClose}>
			<ModalOverlay />
			<ModalContent>
				<ModalHeader>{title}</ModalHeader>
				<ModalCloseButton disabled={isSaving} />
				<ModalBody marginBottom={8}>
					{children}
				</ModalBody>

				<ModalFooter>
					<Button size={'sm'} isDisabled={isSaving} colorScheme='red' variant={'outline'} mr={3} onClick={onClose}>
						{cancelText}
					</Button>
					<Button size={'sm'} display={withSaveButton ? 'inherit' : 'none'} isDisabled={isSaving} variant={'outline'} colorScheme={'green'} onClick={onSave} disabled={isSaving}>{isSaving ? 'Menyimpan...' : saveText}</Button>
				</ModalFooter>
			</ModalContent>
		</Modal>
	)
}

export default MyModal
