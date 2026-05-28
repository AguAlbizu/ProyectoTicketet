// Displays raffle information and the user's current entry status.
// Props:
//   raffle     — { id, name, price_per_chance, status, winner_user }
//   userEntry  — { chances } | null (null if user hasn't bought chances yet)
//   onBuyChances — callback() to open the BuyChancesModal
function RaffleCard({ raffle, userEntry, onBuyChances }) {
  // TODO: render raffle name, price per chance
  // TODO: if status === "pendiente": show user's current chances and a "Comprar chances" button
  // TODO: if status === "realizado": show winner information

  return <div>{/* TODO: implement RaffleCard UI */}</div>
}

export default RaffleCard
