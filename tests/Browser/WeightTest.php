<?php

namespace Tests\Browser;

use App\User;
use App\Weight;
use Illuminate\Foundation\Testing\DatabaseMigrations;
use Laravel\Dusk\Browser;
use Tests\DuskTestCase;

class WeightTest extends DuskTestCase
{
    use DatabaseMigrations;

    public function setUp()
    {
        parent::setUp();
    }

     /** @test */
    public function user_should_see_all_weight_entries()
    {
        $user = factory(User::class)->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Weight Entries')
                    ->clickLink('See all entries')
                    ->assertPathIs('/weight/all')
                    ->assertSee('Weight Entries')
                    ->logout();
        });
    }

    /** @test */
    public function user_can_enter_a_new_weight_entry()
    {
        $user = factory(User::class)->create();
        $this->browse(function (Browser $browser) use ($user) {
            $browser->loginAs($user)
                    ->visit('/')
                    ->assertSee('Latest Weight Entries')
                    ->clickLink('Enter new weight entry')
                    ->assertSee('Add New Weight Entry')
                    ->type('weight', 155.9)
                    ->script([
                        "document.querySelector('#entered-date').value = '2017-12-18'",
                    ]);

            $browser->press('Add Weight Entry')
                    ->assertPathIs('/weight/add')
                    ->assertSee('Weight entry added succesfully.')
                    ->logout();
        });
    }
}
