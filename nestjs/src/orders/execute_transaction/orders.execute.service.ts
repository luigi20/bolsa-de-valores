import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';
import { OrderStatus, OrderType } from '@prisma/client';
import { OrderExecuteDTO } from './order_execute.dto';

@Injectable()
export class OrdersExecuteService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute({
    broker_transaction_id,
    negotiated_shares,
    order_id,
    price,
    related_investor_id,
    status,
  }: OrderExecuteDTO) {
    return this.prismaService.$transaction(async (prisma) => {
      const order = await prisma.order.findUniqueOrThrow({
        where: { id: order_id },
      });
      await prisma.order.update({
        where: { id: order_id },
        data: {
          partial: order.partial - negotiated_shares,
          status: status,
          Transactions: {
            create: {
              broker_transaction_id: broker_transaction_id,
              related_investor_id: related_investor_id,
              shares: negotiated_shares,
              price: price,
            },
          },
        },
      });
      if (status === OrderStatus.CLOSED) {
        await prisma.asset.update({
          where: { id: order.asset_id },
          data: {
            price: price,
          },
        });
        const wallet_asset = await prisma.walletAsset.findUnique({
          where: {
            wallet_id_asset_id: {
              asset_id: order.asset_id,
              wallet_id: order.wallet_id,
            },
          },
        });
        if (wallet_asset) {
          await prisma.walletAsset.update({
            where: {
              wallet_id_asset_id: {
                asset_id: order.asset_id,
                wallet_id: order.wallet_id,
              },
            },
            data: {
              shares:
                order.type === OrderType.BUY
                  ? wallet_asset.shares + order.shares
                  : wallet_asset.shares - order.shares,
            },
          });
        } else {
          await prisma.walletAsset.create({
            data: {
              asset_id: order.asset_id,
              wallet_id: order.wallet_id,
              shares: negotiated_shares,
            },
          });
        }
      }
    });
  }
}
