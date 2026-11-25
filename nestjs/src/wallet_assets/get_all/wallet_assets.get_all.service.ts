import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';

@Injectable()
export class WalletAssetsGetAllService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute(wallet_id: string) {
    return await this.prismaService.walletAsset.findMany({
      where: {
        wallet_id: wallet_id,
      },
      include: {
        asset: {
          select: {
            id: true,
            symbol: true,
            price: true,
          },
        },
      },
    });
  }
}
