<?php

namespace App\Console\Commands;

use App\User;
use App\Weight;
use Carbon\Carbon;
use Illuminate\Console\Command;

class ImportWeightEntries extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'weight:import';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Import initial weight entries';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
        $this->info('Importing predefined readings...');

        $user = User::where('email', '=', 'jonathantorres41@gmail.com')->first();

        if (is_null($user)) {
            $this->error('The user with email "jonathantorres41@gmail.com" was not found');
            return;
        }

        $this->importWeightEntries($user);

        $this->info('Done');
    }

    private function importWeightEntries(User $user)
    {
        Weight::create([
            'user_id' => $user->id,
            'weight' => 148.8,
            'entered_date' => Carbon::createFromDate(2016, 6, 13), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 152.4,
            'entered_date' => Carbon::createFromDate(2016, 6, 27), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 153.6,
            'entered_date' => Carbon::createFromDate(2016, 6, 28), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 155.8,
            'entered_date' => Carbon::createFromDate(2016, 6, 29), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 153.4,
            'entered_date' => Carbon::createFromDate(2016, 6, 30), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 154,
            'entered_date' => Carbon::createFromDate(2016, 7, 1), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 154.8,
            'entered_date' => Carbon::createFromDate(2016, 7, 2), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 155,
            'entered_date' => Carbon::createFromDate(2016, 7, 3), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 154,
            'entered_date' => Carbon::createFromDate(2016, 7, 4), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 154.4,
            'entered_date' => Carbon::createFromDate(2016, 7, 5), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 152.8,
            'entered_date' => Carbon::createFromDate(2016, 7, 6), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 154.8,
            'entered_date' => Carbon::createFromDate(2016, 7, 7), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 154.4,
            'entered_date' => Carbon::createFromDate(2016, 7, 8), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 156,
            'entered_date' => Carbon::createFromDate(2016, 7, 9), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 156,
            'entered_date' => Carbon::createFromDate(2016, 7, 10), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 154.2,
            'entered_date' => Carbon::createFromDate(2016, 7, 11), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 155,
            'entered_date' => Carbon::createFromDate(2016, 7, 12), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 153.8,
            'entered_date' => Carbon::createFromDate(2016, 7, 14), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 154.8,
            'entered_date' => Carbon::createFromDate(2016, 7, 15), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 152.8,
            'entered_date' => Carbon::createFromDate(2016, 7, 17), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 153.8,
            'entered_date' => Carbon::createFromDate(2016, 7, 18), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 156.4,
            'entered_date' => Carbon::createFromDate(2016, 7, 19), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 155.6,
            'entered_date' => Carbon::createFromDate(2016, 7, 21), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157,
            'entered_date' => Carbon::createFromDate(2016, 7, 22), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 156,
            'entered_date' => Carbon::createFromDate(2016, 7, 23), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157.6,
            'entered_date' => Carbon::createFromDate(2016, 7, 24), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157,
            'entered_date' => Carbon::createFromDate(2016, 7, 25), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157.4,
            'entered_date' => Carbon::createFromDate(2016, 7, 26), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 156.4,
            'entered_date' => Carbon::createFromDate(2016, 7, 27), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 156.9,
            'entered_date' => Carbon::createFromDate(2016, 7, 30), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158,
            'entered_date' => Carbon::createFromDate(2016, 7, 31), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158.6,
            'entered_date' => Carbon::createFromDate(2016, 8, 1), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 159,
            'entered_date' => Carbon::createFromDate(2016, 8, 2), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157.6,
            'entered_date' => Carbon::createFromDate(2016, 8, 3), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 156.4,
            'entered_date' => Carbon::createFromDate(2016, 8, 6), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 159,
            'entered_date' => Carbon::createFromDate(2016, 8, 7), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158.6,
            'entered_date' => Carbon::createFromDate(2016, 8, 8), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 159.4,
            'entered_date' => Carbon::createFromDate(2016, 8, 10), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158.6,
            'entered_date' => Carbon::createFromDate(2016, 8, 11), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158.5,
            'entered_date' => Carbon::createFromDate(2016, 8, 13), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157.6,
            'entered_date' => Carbon::createFromDate(2016, 8, 15), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157.6,
            'entered_date' => Carbon::createFromDate(2016, 8, 16), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157.6,
            'entered_date' => Carbon::createFromDate(2016, 8, 17), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 159,
            'entered_date' => Carbon::createFromDate(2016, 8, 20), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158.8,
            'entered_date' => Carbon::createFromDate(2016, 8, 21), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158,
            'entered_date' => Carbon::createFromDate(2016, 8, 22), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158,
            'entered_date' => Carbon::createFromDate(2016, 8, 23), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 159,
            'entered_date' => Carbon::createFromDate(2016, 8, 24), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 161.2,
            'entered_date' => Carbon::createFromDate(2016, 8, 25), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 161,
            'entered_date' => Carbon::createFromDate(2016, 8, 27), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 161,
            'entered_date' => Carbon::createFromDate(2016, 8, 28), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 160,
            'entered_date' => Carbon::createFromDate(2016, 8, 29), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 160,
            'entered_date' => Carbon::createFromDate(2016, 9, 4), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 161.4,
            'entered_date' => Carbon::createFromDate(2016, 9, 5), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 161,
            'entered_date' => Carbon::createFromDate(2016, 9, 8), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 159.8,
            'entered_date' => Carbon::createFromDate(2016, 9, 12), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 161,
            'entered_date' => Carbon::createFromDate(2016, 9, 14), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 160.2,
            'entered_date' => Carbon::createFromDate(2016, 9, 18), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 161.2,
            'entered_date' => Carbon::createFromDate(2016, 9, 19), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 160,
            'entered_date' => Carbon::createFromDate(2016, 9, 25), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158,
            'entered_date' => Carbon::createFromDate(2016, 10, 10), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 169,
            'entered_date' => Carbon::createFromDate(2017, 3, 12), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 170.8,
            'entered_date' => Carbon::createFromDate(2017, 3, 20), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 171.6,
            'entered_date' => Carbon::createFromDate(2017, 3, 23), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 170.2,
            'entered_date' => Carbon::createFromDate(2017, 3, 27), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 168.6,
            'entered_date' => Carbon::createFromDate(2017, 4, 11), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 168.6,
            'entered_date' => Carbon::createFromDate(2017, 4, 17), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 166.8,
            'entered_date' => Carbon::createFromDate(2017, 4, 24), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 166.2,
            'entered_date' => Carbon::createFromDate(2017, 5, 22), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 164.4,
            'entered_date' => Carbon::createFromDate(2017, 6, 4), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 157,
            'entered_date' => Carbon::createFromDate(2017, 8, 3), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 158,
            'entered_date' => Carbon::createFromDate(2017, 8, 14), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 160.6,
            'entered_date' => Carbon::createFromDate(2017, 9, 26), // Year, month, day,
        ]);

        Weight::create([
            'user_id' => $user->id,
            'weight' => 160.6,
            'entered_date' => Carbon::createFromDate(2017, 11, 15), // Year, month, day,
        ]);
    }
}
