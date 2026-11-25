import { OrderType } from '@prisma/client';

export class InitTransactionDTO {
  asset_id: string;
  wallet_id: string;
  shares: number;
  type: OrderType;
  price: number;
}
