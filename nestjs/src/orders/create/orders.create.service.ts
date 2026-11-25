import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';
import { OrderStatus, OrderType } from '@prisma/client';
import { InitTransactionDTO } from './order.dto';

@Injectable()
export class OrdersCreateService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute({
    asset_id,
    price,
    shares,
    type,
    wallet_id,
  }: InitTransactionDTO) {
    return await this.prismaService.order.create({
      data: {
        asset_id: asset_id,
        price: price,
        shares: shares,
        wallet_id: wallet_id,
        type: type,
        status: OrderStatus.PENDING,
        partial: shares,
      },
    });
  }
}
