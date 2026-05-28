// Modal dialog for transferring a ticket to another user.
// Props:
//   ticketId  — ID of the ticket to transfer
//   isOpen    — boolean controlling visibility
//   onClose   — callback to close the modal
//   onConfirm — callback(ticketId, targetEmail) after confirmation
function TransferModal({ ticketId, isOpen, onClose, onConfirm }) {
  // TODO: render an input field for target user email
  // TODO: validate email format before calling onConfirm
  // TODO: show loading state while the transfer API call is in progress

  if (!isOpen) return null

  return <div>{/* TODO: implement TransferModal UI */}</div>
}

export default TransferModal
