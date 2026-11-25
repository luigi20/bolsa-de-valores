import { Test, TestingModule } from '@nestjs/testing';
import { WalletsGetAllController } from './wallets.get_all.controller';
import { WalletsGetAllService } from './wallets.get_all.service';

describe('WalletsController', () => {
  let controller: WalletsGetAllController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [WalletsGetAllController],
      providers: [WalletsGetAllService],
    }).compile();

    controller = module.get<WalletsGetAllController>(WalletsGetAllController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
