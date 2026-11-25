import { Injectable } from '@nestjs/common';
import { PrismaService } from '../../infra/prisma.service';

export interface IWalletAssetsCreateRequest {
  wallet_id: string;
  asset_id: string;
  shares: number;
}
@Injectable()
export class WalletAssetsCreateService {
  constructor(private readonly prismaService: PrismaService) {}

  async execute({ wallet_id, asset_id, shares }: IWalletAssetsCreateRequest) {
    return await this.prismaService.walletAsset.create({
      data: {
        wallet_id: wallet_id,
        asset_id: asset_id,
        shares: shares,
      },
    });
  }
}
