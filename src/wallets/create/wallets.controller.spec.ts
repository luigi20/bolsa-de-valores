import { Test, TestingModule } from '@nestjs/testing';
import { WalletsCreateController } from './wallets.create.controller';
import { WalletsCreateService } from './wallets.create.service';

describe('WalletsController', () => {
  let controller: WalletsCreateController;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [WalletsCreateController],
      providers: [WalletsCreateService],
    }).compile();

    controller = module.get<WalletsCreateController>(WalletsCreateController);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });
});
