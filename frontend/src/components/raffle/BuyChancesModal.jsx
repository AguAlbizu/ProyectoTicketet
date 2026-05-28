// Modal dialog for purchasing raffle chances.
// Props:
//   raffle    — { id, name, price_per_chance }
//   isOpen    — boolean controlling visibility
//   onClose   — callback to close the modal
//   onConfirm — callback(raffleId, quantity) after confirmation
function BuyChancesModal({ raffle, isOpen, onClose, onConfirm }) {
  // TODO: render a numeric input (quantity >= 1)
  // TODO: show total cost = quantity * raffle.price_per_chance
  // TODO: show loading state while API call is in progress
  // TODO: call onConfirm(raffle.id, quantity) on submit

  if (!isOpen) return null

  return <div>{/* TODO: implement BuyChancesModal UI */}</div>
}

export default BuyChancesModal
