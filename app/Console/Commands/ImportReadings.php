<?php

namespace App\Console\Commands;

use Carbon\Carbon;
use Illuminate\Console\Command;
use Illuminate\Support\Facades\DB;

class ImportReadings extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'readings:import';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = 'Import initial blood pressure readings.';

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

        DB::table('blood_pressures')->insert($this->readingsToImport());

        $this->info('Done.');
    }

    /**
     * Define and return the readings to import.
     *
     * @todo  Add the actual readings.
     *
     * @return array
     */
    private function readingsToImport()
    {
        return [
            // [
            //     'systolic' => 123,
            //     'diastolic' => 81,
            //     'pulse' => 75,
            //     'reading_date' => Carbon::createFromDate(2016, 6, 12),
            //     'created_at' => Carbon::now(),
            //     'updated_at' => Carbon::now(),
            // ],
            [
                'systolic' => 145,
                'diastolic' => 74,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 13), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 119,
                'diastolic' => 64,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 25), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 138,
                'diastolic' => 71,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 26), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 133,
                'diastolic' => 72,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 27), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 131,
                'diastolic' => 72,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 29), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 141,
                'diastolic' => 84,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 4, 30), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 135,
                'diastolic' => 64,
                'pulse' => 0,
                'reading_date' => Carbon::createFromDate(2017, 5, 1), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 124,
                'diastolic' => 66,
                'pulse' => 78,
                'reading_date' => Carbon::createFromDate(2017, 5, 2), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 106,
                'diastolic' => 66,
                'pulse' => 106,
                'reading_date' => Carbon::createFromDate(2017, 5, 2), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 129,
                'diastolic' => 64,
                'pulse' => 102,
                'reading_date' => Carbon::createFromDate(2017, 5, 2), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 121,
                'diastolic' => 77,
                'pulse' => 80,
                'reading_date' => Carbon::createFromDate(2017, 5, 6), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 128,
                'diastolic' => 75,
                'pulse' => 70,
                'reading_date' => Carbon::createFromDate(2017, 5, 6), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 127,
                'diastolic' => 71,
                'pulse' => 73,
                'reading_date' => Carbon::createFromDate(2017, 5, 6), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 127,
                'diastolic' => 67,
                'pulse' => 83,
                'reading_date' => Carbon::createFromDate(2017, 5, 7), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 128,
                'diastolic' => 69,
                'pulse' => 81,
                'reading_date' => Carbon::createFromDate(2017, 5, 7), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 122,
                'diastolic' => 67,
                'pulse' => 61,
                'reading_date' => Carbon::createFromDate(2017, 5, 8), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 128,
                'diastolic' => 63,
                'pulse' => 61,
                'reading_date' => Carbon::createFromDate(2017, 5, 8), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 115,
                'diastolic' => 62,
                'pulse' => 67,
                'reading_date' => Carbon::createFromDate(2017, 5, 9), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 115,
                'diastolic' => 63,
                'pulse' => 73,
                'reading_date' => Carbon::createFromDate(2017, 5, 9), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 120,
                'diastolic' => 71,
                'pulse' => 74,
                'reading_date' => Carbon::createFromDate(2017, 5, 11), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 117,
                'diastolic' => 61,
                'pulse' => 74,
                'reading_date' => Carbon::createFromDate(2017, 5, 13), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 129,
                'diastolic' => 65,
                'pulse' => 98,
                'reading_date' => Carbon::createFromDate(2017, 5, 13), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 97,
                'diastolic' => 45,
                'pulse' => 72,
                'reading_date' => Carbon::createFromDate(2017, 5, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 110,
                'diastolic' => 55,
                'pulse' => 90,
                'reading_date' => Carbon::createFromDate(2017, 5, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 108,
                'diastolic' => 59,
                'pulse' => 89,
                'reading_date' => Carbon::createFromDate(2017, 5, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 126,
                'diastolic' => 70,
                'pulse' => 74,
                'reading_date' => Carbon::createFromDate(2017, 5, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 131,
                'diastolic' => 73,
                'pulse' => 74,
                'reading_date' => Carbon::createFromDate(2017, 5, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 114,
                'diastolic' => 52,
                'pulse' => 67,
                'reading_date' => Carbon::createFromDate(2017, 5, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 120,
                'diastolic' => 71,
                'pulse' => 94,
                'reading_date' => Carbon::createFromDate(2017, 5, 15), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 120,
                'diastolic' => 71,
                'pulse' => 96,
                'reading_date' => Carbon::createFromDate(2017, 5, 15), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 122,
                'diastolic' => 66,
                'pulse' => 69,
                'reading_date' => Carbon::createFromDate(2017, 5, 17), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 120,
                'diastolic' => 64,
                'pulse' => 86,
                'reading_date' => Carbon::createFromDate(2017, 5, 17), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 122,
                'diastolic' => 74,
                'pulse' => 88,
                'reading_date' => Carbon::createFromDate(2017, 5, 18), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 118,
                'diastolic' => 70,
                'pulse' => 92,
                'reading_date' => Carbon::createFromDate(2017, 5, 18), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 132,
                'diastolic' => 72,
                'pulse' => 85,
                'reading_date' => Carbon::createFromDate(2017, 5, 19), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 121,
                'diastolic' => 70,
                'pulse' => 80,
                'reading_date' => Carbon::createFromDate(2017, 5, 19), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 131,
                'diastolic' => 63,
                'pulse' => 72,
                'reading_date' => Carbon::createFromDate(2017, 5, 20), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 124,
                'diastolic' => 63,
                'pulse' => 69,
                'reading_date' => Carbon::createFromDate(2017, 5, 20), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 125,
                'diastolic' => 68,
                'pulse' => 71,
                'reading_date' => Carbon::createFromDate(2017, 5, 21), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 119,
                'diastolic' => 69,
                'pulse' => 66,
                'reading_date' => Carbon::createFromDate(2017, 5, 21), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 118,
                'diastolic' => 64,
                'pulse' => 80,
                'reading_date' => Carbon::createFromDate(2017, 5, 23), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 112,
                'diastolic' => 61,
                'pulse' => 78,
                'reading_date' => Carbon::createFromDate(2017, 5, 23), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 114,
                'diastolic' => 62,
                'pulse' => 101,
                'reading_date' => Carbon::createFromDate(2017, 5, 26), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 107,
                'diastolic' => 60,
                'pulse' => 96,
                'reading_date' => Carbon::createFromDate(2017, 5, 26), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 118,
                'diastolic' => 63,
                'pulse' => 98,
                'reading_date' => Carbon::createFromDate(2017, 5, 26), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 146,
                'diastolic' => 72,
                'pulse' => 90,
                'reading_date' => Carbon::createFromDate(2017, 5, 31), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 139,
                'diastolic' => 68,
                'pulse' => 95,
                'reading_date' => Carbon::createFromDate(2017, 5, 31), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 122,
                'diastolic' => 72,
                'pulse' => 86,
                'reading_date' => Carbon::createFromDate(2017, 5, 31), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 115,
                'diastolic' => 73,
                'pulse' => 84,
                'reading_date' => Carbon::createFromDate(2017, 5, 31), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 131,
                'diastolic' => 64,
                'pulse' => 86,
                'reading_date' => Carbon::createFromDate(2017, 6, 1), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 131,
                'diastolic' => 64,
                'pulse' => 87,
                'reading_date' => Carbon::createFromDate(2017, 6, 1), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 130,
                'diastolic' => 69,
                'pulse' => 70,
                'reading_date' => Carbon::createFromDate(2017, 6, 6), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 130,
                'diastolic' => 71,
                'pulse' => 71,
                'reading_date' => Carbon::createFromDate(2017, 6, 6), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 122,
                'diastolic' => 65,
                'pulse' => 67,
                'reading_date' => Carbon::createFromDate(2017, 6, 7), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 114,
                'diastolic' => 63,
                'pulse' => 74,
                'reading_date' => Carbon::createFromDate(2017, 6, 7), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 128,
                'diastolic' => 67,
                'pulse' => 61,
                'reading_date' => Carbon::createFromDate(2017, 6, 11), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 126,
                'diastolic' => 66,
                'pulse' => 58,
                'reading_date' => Carbon::createFromDate(2017, 6, 11), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 131,
                'diastolic' => 69,
                'pulse' => 95,
                'reading_date' => Carbon::createFromDate(2017, 6, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 128,
                'diastolic' => 67,
                'pulse' => 96,
                'reading_date' => Carbon::createFromDate(2017, 6, 14), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 120,
                'diastolic' => 61,
                'pulse' => 76,
                'reading_date' => Carbon::createFromDate(2017, 6, 15), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 125,
                'diastolic' => 63,
                'pulse' => 95,
                'reading_date' => Carbon::createFromDate(2017, 6, 17), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 126,
                'diastolic' => 72,
                'pulse' => 88,
                'reading_date' => Carbon::createFromDate(2017, 6, 22), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 132,
                'diastolic' => 74,
                'pulse' => 61,
                'reading_date' => Carbon::createFromDate(2017, 6, 25), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 120,
                'diastolic' => 63,
                'pulse' => 95,
                'reading_date' => Carbon::createFromDate(2017, 6, 27), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 121,
                'diastolic' => 68,
                'pulse' => 80,
                'reading_date' => Carbon::createFromDate(2017, 7, 4), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 127,
                'diastolic' => 70,
                'pulse' => 73,
                'reading_date' => Carbon::createFromDate(2017, 7, 7), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
            [
                'systolic' => 124,
                'diastolic' => 66,
                'pulse' => 68,
                'reading_date' => Carbon::createFromDate(2017, 7, 9), // Year, month, day
                'created_at' => Carbon::now(),
                'updated_at' => Carbon::now(),
            ],
        ];
    }
}
