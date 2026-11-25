export class OrderExecuteDTO {
  order_id: string;
  wallet_id: string;
  status: 'OPEN' | 'CLOSED';
  related_investor_id: string;
  broker_transaction_id: string;
  negotiated_shares: number;
  price: number;
}
