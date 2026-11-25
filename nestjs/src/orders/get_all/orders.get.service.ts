import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';

@Injectable()
export class OrdersGetService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute(wallet_id: string) {
    return await this.prismaService.order.findMany({
      where: {
        wallet_id: wallet_id,
      },
      include: {
        transactions: true,
        asset: {
          select: {
            id: true,
            symbol: true,
          },
        },
      },
      orderBy: {
        updated_at: 'desc',
      },
    });
  }
}
